# VkLatencyTimingsFrameReportNV(3) Manual Page

## Name

VkLatencyTimingsFrameReportNV - Structure containing latency data



## [](#_c_specification)C Specification

The [VkLatencyTimingsFrameReportNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkLatencyTimingsFrameReportNV.html) structure describes latency data returned by [vkGetLatencyTimingsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetLatencyTimingsNV.html)

```c++
// Provided by VK_NV_low_latency2
typedef struct VkLatencyTimingsFrameReportNV {
    VkStructureType    sType;
    const void*        pNext;
    uint64_t           presentID;
    uint64_t           inputSampleTimeUs;
    uint64_t           simStartTimeUs;
    uint64_t           simEndTimeUs;
    uint64_t           renderSubmitStartTimeUs;
    uint64_t           renderSubmitEndTimeUs;
    uint64_t           presentStartTimeUs;
    uint64_t           presentEndTimeUs;
    uint64_t           driverStartTimeUs;
    uint64_t           driverEndTimeUs;
    uint64_t           osRenderQueueStartTimeUs;
    uint64_t           osRenderQueueEndTimeUs;
    uint64_t           gpuRenderStartTimeUs;
    uint64_t           gpuRenderEndTimeUs;
} VkLatencyTimingsFrameReportNV;
```

## [](#_members)Members

The members of the [VkLatencyTimingsFrameReportNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkLatencyTimingsFrameReportNV.html) structure describe the following:

## [](#_description)Description

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `presentID` is the application provided value that is used to associate the timestamp with a `vkQueuePresentKHR` command using [VkPresentIdKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentIdKHR.html)::`pPresentIds` or [VkPresentId2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentId2KHR.html)::`pPresentIds` for a given present.
- `simStartTimeUs` is the timestamp written when `vkSetLatencyMarkerNV` is called with the `VkLatencyMarkerNV` enum `VK_LATENCY_MARKER_SIMULATION_START_NV`.
- `simEndTimeUs` is the timestamp written when `vkSetLatencyMarkerNV` is called with the `VkLatencyMarkerNV` enum `VK_LATENCY_MARKER_SIMULATION_END_NV`
- `renderStartTimeUs` is the timestamp written when `vkSetLatencyMarkerNV` is called with the `VkLatencyMarkerNV` enum `VK_LATENCY_MARKER_RENDERSUBMIT_START_NV`.
- `renderEndTimeUs` is the timestamp written when `vkSetLatencyMarkerNV` is called with the `VkLatencyMarkerNV` enum `VK_LATENCY_MARKER_RENDERSUBMIT_END_NV`.
- `presentStartTimeUs` is the timestamp written when `vkSetLatencyMarkerNV` is called with the `VkLatencyMarkerNV` enum `VK_LATENCY_MARKER_PRESENT_START_NV`.
- `presentEndTimeUs` is the timestamp written when `vkSetLatencyMarkerNV` is called with the `VkLatencyMarkerNV` enum `VK_LATENCY_MARKER_PRESENT_END_NV`.
- `driverStartTimeUs` is the timestamp written when the first `vkQueueSubmit` for the frame is called.
- `driverEndTimeUs` is the timestamp written when the final `vkQueueSubmit` hands off from the Vulkan Driver.
- `osRenderQueueStartTimeUs` is the timestamp written when the final `vkQueueSubmit` hands off from the Vulkan Driver.
- `osRenderQueueEndTimeUs` is the timestamp written when the first submission reaches the GPU.
- `gpuRenderStartTimeUs` is the timestamp written when the first submission reaches the GPU.
- `gpuRenderEndTimeUs` is the timestamp written when the final submission finishes on the GPU for the frame.

Valid Usage (Implicit)

- [](#VUID-VkLatencyTimingsFrameReportNV-sType-sType)VUID-VkLatencyTimingsFrameReportNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_LATENCY_TIMINGS_FRAME_REPORT_NV`

## [](#_see_also)See Also

[VK\_NV\_low\_latency2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_low_latency2.html), [VkGetLatencyMarkerInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGetLatencyMarkerInfoNV.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkLatencyTimingsFrameReportNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0