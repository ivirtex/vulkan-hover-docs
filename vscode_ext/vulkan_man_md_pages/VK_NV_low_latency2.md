# VK_NV_low_latency2(3) Manual Page

## Name

VK_NV_low_latency2 - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

506

## <a href="#_revision" class="anchor"></a>Revision

2

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[Version 1.2](#versions-1.2)  
or  
[VK_KHR_timeline_semaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_timeline_semaphore.html)  

## <a href="#_contact" class="anchor"></a>Contact

- Charles Hansen <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_low_latency2%5D%20@cshansen%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_low_latency2%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>cshansen</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2023-09-25

**Contributors**  
- Charles Hansen, NVIDIA

- Liam Middlebrook, NVIDIA

- Lionel Duc, NVIDIA

- James Jones, NVIDIA

- Eric Sullivan, NVIDIA

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkGetLatencyTimingsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetLatencyTimingsNV.html)

- [vkLatencySleepNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkLatencySleepNV.html)

- [vkQueueNotifyOutOfBandNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkQueueNotifyOutOfBandNV.html)

- [vkSetLatencyMarkerNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkSetLatencyMarkerNV.html)

- [vkSetLatencySleepModeNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkSetLatencySleepModeNV.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkGetLatencyMarkerInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGetLatencyMarkerInfoNV.html)

- [VkLatencySleepInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLatencySleepInfoNV.html)

- [VkLatencySleepModeInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLatencySleepModeInfoNV.html)

- [VkLatencyTimingsFrameReportNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLatencyTimingsFrameReportNV.html)

- [VkOutOfBandQueueTypeInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOutOfBandQueueTypeInfoNV.html)

- [VkSetLatencyMarkerInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSetLatencyMarkerInfoNV.html)

- Extending [VkSubmitInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubmitInfo.html),
  [VkSubmitInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubmitInfo2.html):

  - [VkLatencySubmissionPresentIdNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLatencySubmissionPresentIdNV.html)

- Extending [VkSurfaceCapabilities2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceCapabilities2KHR.html):

  - [VkLatencySurfaceCapabilitiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLatencySurfaceCapabilitiesNV.html)

- Extending [VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainCreateInfoKHR.html):

  - [VkSwapchainLatencyCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainLatencyCreateInfoNV.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkLatencyMarkerNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkLatencyMarkerNV.html)

- [VkOutOfBandQueueTypeNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOutOfBandQueueTypeNV.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_NV_LOW_LATENCY_2_EXTENSION_NAME`

- `VK_NV_LOW_LATENCY_2_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_GET_LATENCY_MARKER_INFO_NV`

  - `VK_STRUCTURE_TYPE_LATENCY_SLEEP_INFO_NV`

  - `VK_STRUCTURE_TYPE_LATENCY_SLEEP_MODE_INFO_NV`

  - `VK_STRUCTURE_TYPE_LATENCY_SUBMISSION_PRESENT_ID_NV`

  - `VK_STRUCTURE_TYPE_LATENCY_SURFACE_CAPABILITIES_NV`

  - `VK_STRUCTURE_TYPE_LATENCY_TIMINGS_FRAME_REPORT_NV`

  - `VK_STRUCTURE_TYPE_OUT_OF_BAND_QUEUE_TYPE_INFO_NV`

  - `VK_STRUCTURE_TYPE_SET_LATENCY_MARKER_INFO_NV`

  - `VK_STRUCTURE_TYPE_SWAPCHAIN_LATENCY_CREATE_INFO_NV`

## <a href="#_description" class="anchor"></a>Description

This extension gives applications timing suggestions on when to start
the recording of new frames to reduce the latency between input sampling
and frame presentation. Applications can accomplish this through the
extension by calling
[vkSetLatencySleepModeNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkSetLatencySleepModeNV.html) to allow the
driver to pace a given swapchain, then calling
[vkLatencySleepNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkLatencySleepNV.html) before input sampling to delay
the start of the CPU side work. Additional methods and structures are
provided to give insight into the latency pipeline of an application
through the latency markers.
[`VK_NV_low_latency`](VK_NV_low_latency.html) provides legacy support
for applications that make use of the NVIDIA Reflex SDK whereas new
implementations should use the
[`VK_NV_low_latency2`](VK_NV_low_latency2.html) extension.

## <a href="#_issues" class="anchor"></a>Issues

1\) How does Low Latency 2 work with applications that utilize device
groups?

Low Latency 2 does not support device groups.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 2, 2023-11-15 (Charles Hansen)

  - Update vkGetLatencyTimingsNV. This is a breaking API change which
    brings behavior in line with other array querying commands. More
    background can be found in
    <a href="https://github.com/KhronosGroup/Vulkan-Docs/issues/2269"
    class="bare">https://github.com/KhronosGroup/Vulkan-Docs/issues/2269</a>

- Revision 1, 2023-09-25 (Charles Hansen)

  - Internal revisions

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_NV_low_latency2"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
