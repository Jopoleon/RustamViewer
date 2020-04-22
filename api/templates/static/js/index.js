function getVars() {
    var VARS = null;
    $.ajax({
        type: "get",
        url: "/vars",
        success: function (data, textStatus, jqXHR) {
            // console.log(data);
            VARS = data;
        },
    });
    return VARS
}

function getArs() {
    var VARS = null;
    $.ajax({
        type: "get",
        url: "/ars",
        success: function (data, textStatus, jqXHR) {
            // console.log(data);
            VARS = data;
        },
    });
    return VARS
}

function getCallsAll() {
    var VARS = null;
    $.ajax({
        type: "get",
        url: "/calls",
        success: function (data, textStatus, jqXHR) {
            // console.log(data);
            VARS = data;
        },
    });
    return VARS
}

function getCallsOut() {
    var VARS = null;
    $.ajax({
        type: "get",
        url: "/callsout",
        success: function (data, textStatus, jqXHR) {
            // console.log(data);
            VARS = data;
        },
    });
    return VARS
}

$(document).ready(function () {
    $("#btn-grp > button").on("click", function () {
        var defaultClass = "btn btn-secondary";
        var toBeAssignedClass = $(this).attr("data-btn-class");

        $("#btn-grp > button").attr("class", defaultClass);
        $(this).attr("class", toBeAssignedClass);
    });
    $('#wrapperTableCalls').hide();
    $('#wrapperTableCallsOut').hide();
    $('#wrapperTableARS').show();
    $('#wrapperTableVARS').hide();

    $('#btnARS').click(function () {
        $('#wrapperTableCalls').hide();
        $('#wrapperTableCallsOut').hide();
        $('#wrapperTableARS').show();
        $('#wrapperTableVARS').hide();
    });
    $('#btnVARS').click(function () {
        $('#wrapperTableCalls').hide();
        $('#wrapperTableCallsOut').hide();
        $('#wrapperTableARS').hide();
        $('#wrapperTableVARS').show();
    });
    $('#btnCallsOut').click(function () {
        $('#wrapperTableCalls').hide();
        $('#wrapperTableCallsOut').show();
        $('#wrapperTableARS').hide();
        $('#wrapperTableVARS').hide();
    });
    $('#btnCallsAll').click(function () {
        $('#wrapperTableCalls').show();
        $('#wrapperTableCallsOut').hide();
        $('#wrapperTableARS').hide();
        $('#wrapperTableVARS').hide();
    });

    $('#tableVARS').DataTable({
        dataSrc: "",
        ajax: {
            url: '/vars',
            dataSrc: ''
        },
        order: [[1, 'asc']],
        columns: [
            {
                "data": "id", // can be null or undefined
                "defaultContent": "<i>Empty</i>"

            },
            {"data": "external_id"},
            {
                "data": "menu_name",
                "defaultContent": "empty"
            },
            {
                "data": "project_id",
                "defaultContent": "empty"
            },
            {
                "data": "ani",
                "defaultContent": "empty"
            },
            {
                "data": "callid",
                "defaultContent": "empty"
            },
            {
                "data": "seq",
                "defaultContent": "empty"
            },
            {
                "data": "action_status",
                "defaultContent": "empty"
            },
            {
                "data": "action_description",
                "defaultContent": "empty"
            },
            {
                "data": "enter_menu_time",
                "defaultContent": "empty"
            },
            {
                "data": "leave_menu_time",
                "defaultContent": "empty"
            },
            {
                "data": "action_time",
                "defaultContent": "empty"
            },
        ],

    });
    $('#tableASR').DataTable({
        dataSrc: "",
        ajax: {
            url: '/ars',
            dataSrc: ''
        },
        order: [[1, 'asc']],
        columns: [
            {
                "data": "menu_name",
                "defaultContent": "empty"
            },
            {
                "data": "project_id",
                "defaultContent": "empty"
            },
            {
                "data": "ani",
                "defaultContent": "empty"
            },
            {
                "data": "callid",
                "defaultContent": "empty"
            },
            {
                "data": "seq",
                "defaultContent": "empty"
            },
            {
                "data": "utterance",
                "defaultContent": "empty"
            },
            {
                "data": "interpretation",
                "defaultContent": "empty"
            },
            {
                "data": "confidence",
                "defaultContent": "empty"
            },
            {
                "data": "inputmode",
                "defaultContent": "empty"
            },
            {
                "data": "waverecord",
                "render": function (data, type, row, meta) {
                    elem = '<audio controls preload="none" class="audioPlayer"><source src="/waverecord/' + row.id + '" type="audio/wav"></audio>'
                    return elem
                }
            },
        ],
    });
    $('#tableCallsAll').DataTable({
        dataSrc: "",
        ajax: {
            url: '/calls',
            dataSrc: ''
        },
        order: [[1, 'asc']],
        columns: [
            {
                "data": "id",
                "defaultContent": "empty"
            },
            {
                "data": "interaction_id",
                "defaultContent": "empty"
            },
            {
                "data": "source_address",
                "defaultContent": "empty"
            },
            {
                "data": "target_address",
                "defaultContent": "empty"
            },
            {
                "data": "interaction_type",
                "defaultContent": "empty"
            },
            {
                "data": "media_type",
                "defaultContent": "empty"
            },
            {
                "data": "start_time",
                "defaultContent": "empty"
            },
            {
                "data": "end_time",
                "defaultContent": "empty"
            },
            {
                "data": "project_id",
                "defaultContent": "empty"
            },
            {
                "data": "customer_data",
                "defaultContent": "empty"
            }, {
                "data": "callid",
                "defaultContent": "empty"
            },
            // {
            //     "data": "profilename",
            //     "defaultContent": "empty"
            // },
            {
                "data": "recording_file_id",
                "defaultContent": "empty"
            },
            {
                "data": "end_time_original",
                "defaultContent": "empty"
            },
            {
                "data": "updated_flag",
                "defaultContent": "empty"
            },
            {
                "data": "create_audit_key",
                "defaultContent": "empty"
            },
            {
                "data": "id",
                "render": function (data, type, row, meta) {
                    elem = '<div class="btn-toolbar" role="group"><button type="button" onclick="Download(' + row.id + ',txt)" class="btn btn-primary"><span class="glyphicon glyphicon-save-file"> </span> TXT </button><button type="button" onclick="Download(' + row.id + ',wav)" class="btn btn-primary"><span class="glyphicon glyphicon-headphones"> </span> WAV</button></div>'
                    return elem
                },
            },

        ],
    });
    $('#tableCallsOutbound').DataTable({
        dataSrc: "",
        ajax: {
            url: '/callsout',
            dataSrc: ''
        },
        order: [[1, 'asc']],
        columns: [
            {
                "data": "id",
                "defaultContent": "empty"
            },
            {
                "data": "contact_attempt_fact_key",
                "defaultContent": "empty"
            },
            {
                "data": "contact_info",
                "defaultContent": "empty"
            },
            {
                "data": "media_type",
                "defaultContent": "empty"
            },
            {
                "data": "dialing_mode",
                "defaultContent": "empty"
            },
            {
                "data": "campaing",
                "defaultContent": "empty"
            },
            {
                "data": "call_result",
                "defaultContent": "empty"
            },
            {
                "data": "record_type",
                "defaultContent": "empty"
            },
            {
                "data": "record_status",
                "defaultContent": "empty"
            },
            {
                "data": "calling_list",
                "defaultContent": "empty"
            },
            {
                "data": "contact_info_type",
                "defaultContent": "empty"
            },
            {
                "data": "time_zone",
                "defaultContent": "empty"
            },
            {
                "data": "callid",
                "defaultContent": "empty"
            },
            {
                "data": "start_time",
                "defaultContent": "empty"
            },
            {
                "data": "end_time",
                "defaultContent": "empty"
            },
            {
                "data": "record_id",
                "defaultContent": "empty"
            },
            {
                "data": "chain_id",
                "defaultContent": "empty"
            },
            {
                "data": "chain_n",
                "defaultContent": "empty"
            },
            {
                "data": "attempt",
                "defaultContent": "empty"
            }, {
                "data": "daily_from",
                "defaultContent": "empty"
            },
            {
                "data": "daily_till",
                "defaultContent": "empty"
            },
            {
                "data": "dial_sched_time",
                "defaultContent": "empty"
            },
            {
                "data": "project_id",
                "defaultContent": "empty"
            },
            {
                "data": "customer_data",
                "defaultContent": "empty"
            },

            // {
            //     "data": "id",
            //     "render": function (data, type, row, meta) {
            //         elem = '<div class="btn-toolbar" role="group"><button type="button" onclick="Download(' + row.id + ',txt)" class="btn btn-primary"><span class="glyphicon glyphicon-save-file"> </span> TXT </button><button type="button" onclick="Download(' + row.id + ',wav)" class="btn btn-primary"><span class="glyphicon glyphicon-headphones"> </span> WAV</button></div>'
            //         return elem
            //     },
            // },

        ],
    });
});

function Download(callID, fileType) {
    console.log(callID);
    console.log(fileType);
    $('#getFileError').empty();

    $.ajax({
        type: "GET",
        dataType: "json",
        url: "/file" + "?fileType=" + fileType + "&callID=" + callID,
        fail: function (xhr, textStatus, errorThrown) {
            $('#getFileError').empty();
            $('#getFileErrorr').append(xhr.responseText);
        },
        error: function (jqXHR, textStatus, errorThrown) {
            $('#getFileError').empty();
            $('#getFileError').append(jqXHR.responseText);
        },
        success: function (data, textStatus, jqXHR) {
            $('#getFileError').empty();

            if (data.type === "txt") {
                downloadTxt(data.name + "." + data.type, atou(data.data), 'data:text/plain; charset=UTF-8,');
            }
            if (data.type === "wav") {
                downloadWav(data.name + "." + data.type, data.data, 'data:audio/ogg;base64,');
            }
        },
    });
}

function atou(b64) {
    return decodeURIComponent(escape(atob(b64)));
}

function downloadWav(filename, bytes, attribute) {
    var element = document.createElement('a');
    element.setAttribute('href', attribute + bytes);
    element.setAttribute('download', filename);

    element.style.display = 'none';
    document.body.appendChild(element);

    element.click();

    document.body.removeChild(element);
}

function downloadTxt(filename, text, attribute) {
    var element = document.createElement('a');
    element.setAttribute('href', attribute + encodeURIComponent(text));
    element.setAttribute('download', filename);

    element.style.display = 'none';
    document.body.appendChild(element);

    element.click();

    document.body.removeChild(element);
}

function saveRecording(recordingName) {
    var a = document.createElement("a");
    document.body.appendChild(a);
    a.style = "display: none";
    var blob = FWRecorder.getBlob(recordingName),
        url = window.URL.createObjectURL(blob);
    a.href = url;
    a.download = recordingName + '-fwr-recording.wav';
    a.click();
    window.URL.revokeObjectURL(url);
    document.body.removeChild(a);
}

function submitFilters() {
    $.ajax({
        type: "get",
        url: "/filterTable",
        data: {
            dnis: $('#dnis').val(),
            ani: $('#ani').val(),
            profileName: $('#profileName').val(),
            utterance: $('#utterance').val(),
            confidence: $('#confidence').val(),
            confidenceType: $('#sign').val(),
        },
        fail: function (xhr, textStatus, errorThrown) {
            $('.submitFiltersError').append(xhr.responseText);
        },
        error: function (jqXHR, textStatus, errorThrown) {
            $('.submitFiltersError').append(jqXHR.responseText);
        },
        success: function (data, textStatus, jqXHR) {
            $('#tableFilter').empty();
            $('#tableFilter').append(data);
        },
    });
}