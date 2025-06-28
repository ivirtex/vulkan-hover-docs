# VkLatencyMarkerNV(3) Manual Page

## Name

VkLatencyMarkerNV - Structure used to mark different points in latency



## [](#_c_specification)C Specification

The [VkLatencyMarkerNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkLatencyMarkerNV.html) enum is defined as:

```c++
// Provided by VK_NV_low_latency2
typedef enum VkLatencyMarkerNV {
    VK_LATENCY_MARKER_SIMULATION_START_NV = 0,
    VK_LATENCY_MARKER_SIMULATION_END_NV = 1,
    VK_LATENCY_MARKER_RENDERSUBMIT_START_NV = 2,
    VK_LATENCY_MARKER_RENDERSUBMIT_END_NV = 3,
    VK_LATENCY_MARKER_PRESENT_START_NV = 4,
    VK_LATENCY_MARKER_PRESENT_END_NV = 5,
    VK_LATENCY_MARKER_INPUT_SAMPLE_NV = 6,
    VK_LATENCY_MARKER_TRIGGER_FLASH_NV = 7,
    VK_LATENCY_MARKER_OUT_OF_BAND_RENDERSUBMIT_START_NV = 8,
    VK_LATENCY_MARKER_OUT_OF_BAND_RENDERSUBMIT_END_NV = 9,
    VK_LATENCY_MARKER_OUT_OF_BAND_PRESENT_START_NV = 10,
    VK_LATENCY_MARKER_OUT_OF_BAND_PRESENT_END_NV = 11,
} VkLatencyMarkerNV;
```

## [](#_description)Description

The members of the [VkLatencyMarkerNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkLatencyMarkerNV.html) are used as arguments for [vkSetLatencyMarkerNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkSetLatencyMarkerNV.html) in the use cases described below:

- `VK_LATENCY_MARKER_SIMULATION_START_NV` **should** be called at the start of the simulation execution each frame, but after the call to `vkLatencySleepNV`.
- `VK_LATENCY_MARKER_SIMULATION_END_NV` **should** be called at the end of the simulation execution each frame.
- `VK_LATENCY_MARKER_RENDERSUBMIT_START_NV` **should** be called at the beginning of the render submission execution each frame. This **should** be wherever Vulkan API calls are made and **must** not span into asynchronous rendering.
- `VK_LATENCY_MARKER_RENDERSUBMIT_END_NV` **should** be called at the end of the render submission execution each frame.
- `VK_LATENCY_MARKER_PRESENT_START_NV` **should** be called just before `vkQueuePresentKHR`.
- `VK_LATENCY_MARKER_PRESENT_END_NV` **should** be called when `vkQueuePresentKHR` returns.
- `VK_LATENCY_MARKER_INPUT_SAMPLE_NV` **should** be called just before the application gathers input data.
- `VK_LATENCY_MARKER_TRIGGER_FLASH_NV` **should** be called anywhere between `VK_LATENCY_MARKER_SIMULATION_START_NV` and `VK_LATENCY_MARKER_SIMULATION_END_NV` whenever a left mouse click occurs.

## [](#_see_also)See Also

[VK\_NV\_low\_latency2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_low_latency2.html), [VkSetLatencyMarkerInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSetLatencyMarkerInfoNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkLatencyMarkerNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0