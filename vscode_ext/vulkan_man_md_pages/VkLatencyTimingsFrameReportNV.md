# VkLatencyTimingsFrameReportNV(3) Manual Page

## Name

VkLatencyTimingsFrameReportNV - Structure containing latency data



## <a href="#_c_specification" class="anchor"></a>C Specification

The [VkLatencyTimingsFrameReportNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLatencyTimingsFrameReportNV.html)
structure describes latency data returned by
[vkGetLatencyTimingsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetLatencyTimingsNV.html)

``` c
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

## <a href="#_members" class="anchor"></a>Members

The members of the
[VkLatencyTimingsFrameReportNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLatencyTimingsFrameReportNV.html)
structure describe the following:

## <a href="#_description" class="anchor"></a>Description

- `presentId` is the application provided value that is used to
  associate the timestamp with a `vkQueuePresentKHR` command using
  [VkPresentIdKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentIdKHR.html)::`pPresentIds` for a given
  present.

- `simStartTimeUs` is the timestamp written when `vkSetLatencyMarkerNV`
  is called with the `VkLatencyMarkerNV` enum
  `VK_LATENCY_MARKER_SIMULATION_START_NV`.

- `simEndTimeUs` is the timestamp written when `vkSetLatencyMarkerNV` is
  called with the `VkLatencyMarkerNV` enum
  `VK_LATENCY_MARKER_SIMULATION_END_NV`

- `renderStartTimeUs` is the timestamp written when
  `vkSetLatencyMarkerNV` is called with the `VkLatencyMarkerNV` enum
  `VK_LATENCY_MARKER_RENDERSUBMIT_START_NV`.

- `renderEndTimeUs` is the timestamp written when `vkSetLatencyMarkerNV`
  is called with the `VkLatencyMarkerNV` enum
  `VK_LATENCY_MARKER_RENDERSUBMIT_END_NV`.

- `presentStartTimeUs` is the timestamp written when
  `vkSetLatencyMarkerNV` is called with the `VkLatencyMarkerNV` enum
  `VK_LATENCY_MARKER_PRESENT_START_NV`.

- `presentEndTimeUs` is the timestamp written when
  `vkSetLatencyMarkerNV` is called with the `VkLatencyMarkerNV` enum
  `VK_LATENCY_MARKER_PRESENT_END_NV`.

- `driverStartTimeUs` is the timestamp written when the first
  `vkQueueSubmit` for the frame is called.

- `driverEndTimeUs` is the timestamp written when the final
  `vkQueueSubmit` hands off from the Vulkan Driver.

- `osRenderQueueStartTimeUs` is the timestamp written when the final
  `vkQueueSubmit` hands off from the Vulkan Driver.

- `osRenderQueueEndTimeUs` is the timestamp written when the first
  submission reaches the GPU.

- `gpuRenderStartTimeUs` is the timestamp written when the first
  submission reaches the GPU.

- `gpuRenderEndTimeUs` is the timestamp written when the final
  submission finishes on the GPU for the frame.

Valid Usage (Implicit)

- <a href="#VUID-VkLatencyTimingsFrameReportNV-sType-sType"
  id="VUID-VkLatencyTimingsFrameReportNV-sType-sType"></a>
  VUID-VkLatencyTimingsFrameReportNV-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_LATENCY_TIMINGS_FRAME_REPORT_NV`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_low_latency2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_low_latency2.html),
[VkGetLatencyMarkerInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGetLatencyMarkerInfoNV.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkLatencyTimingsFrameReportNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
