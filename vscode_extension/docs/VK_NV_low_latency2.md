# VK\_NV\_low\_latency2(3) Manual Page

## Name

VK\_NV\_low\_latency2 - device extension



## [](#_registered_extension_number)Registered Extension Number

506

## [](#_revision)Revision

2

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

     [Vulkan Version 1.2](#versions-1.2)  
     or  
     [VK\_KHR\_timeline\_semaphore](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_timeline_semaphore.html)  
and  
     [VK\_KHR\_present\_id](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_present_id.html)  
     or  
     [VK\_KHR\_present\_id2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_present_id2.html)

## [](#_contact)Contact

- Charles Hansen [\[GitHub\]cshansen](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_low_latency2%5D%20%40cshansen%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_low_latency2%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2023-09-25

**Contributors**

- Charles Hansen, NVIDIA
- Liam Middlebrook, NVIDIA
- Lionel Duc, NVIDIA
- James Jones, NVIDIA
- Eric Sullivan, NVIDIA

## [](#_new_commands)New Commands

- [vkGetLatencyTimingsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetLatencyTimingsNV.html)
- [vkLatencySleepNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkLatencySleepNV.html)
- [vkQueueNotifyOutOfBandNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkQueueNotifyOutOfBandNV.html)
- [vkSetLatencyMarkerNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkSetLatencyMarkerNV.html)
- [vkSetLatencySleepModeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkSetLatencySleepModeNV.html)

## [](#_new_structures)New Structures

- [VkGetLatencyMarkerInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGetLatencyMarkerInfoNV.html)
- [VkLatencySleepInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkLatencySleepInfoNV.html)
- [VkLatencySleepModeInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkLatencySleepModeInfoNV.html)
- [VkLatencyTimingsFrameReportNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkLatencyTimingsFrameReportNV.html)
- [VkOutOfBandQueueTypeInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOutOfBandQueueTypeInfoNV.html)
- [VkSetLatencyMarkerInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSetLatencyMarkerInfoNV.html)
- Extending [VkSubmitInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubmitInfo.html), [VkSubmitInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubmitInfo2.html):
  
  - [VkLatencySubmissionPresentIdNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkLatencySubmissionPresentIdNV.html)
- Extending [VkSurfaceCapabilities2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceCapabilities2KHR.html):
  
  - [VkLatencySurfaceCapabilitiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkLatencySurfaceCapabilitiesNV.html)
- Extending [VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainCreateInfoKHR.html):
  
  - [VkSwapchainLatencyCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainLatencyCreateInfoNV.html)

## [](#_new_enums)New Enums

- [VkLatencyMarkerNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkLatencyMarkerNV.html)
- [VkOutOfBandQueueTypeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOutOfBandQueueTypeNV.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_NV_LOW_LATENCY_2_EXTENSION_NAME`
- `VK_NV_LOW_LATENCY_2_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_GET_LATENCY_MARKER_INFO_NV`
  - `VK_STRUCTURE_TYPE_LATENCY_SLEEP_INFO_NV`
  - `VK_STRUCTURE_TYPE_LATENCY_SLEEP_MODE_INFO_NV`
  - `VK_STRUCTURE_TYPE_LATENCY_SUBMISSION_PRESENT_ID_NV`
  - `VK_STRUCTURE_TYPE_LATENCY_SURFACE_CAPABILITIES_NV`
  - `VK_STRUCTURE_TYPE_LATENCY_TIMINGS_FRAME_REPORT_NV`
  - `VK_STRUCTURE_TYPE_OUT_OF_BAND_QUEUE_TYPE_INFO_NV`
  - `VK_STRUCTURE_TYPE_SET_LATENCY_MARKER_INFO_NV`
  - `VK_STRUCTURE_TYPE_SWAPCHAIN_LATENCY_CREATE_INFO_NV`

## [](#_description)Description

This extension gives applications timing suggestions on when to start the recording of new frames to reduce the latency between input sampling and frame presentation. Applications can accomplish this through the extension by calling [vkSetLatencySleepModeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkSetLatencySleepModeNV.html) to allow the driver to pace a given swapchain, then calling [vkLatencySleepNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkLatencySleepNV.html) before input sampling to delay the start of the CPU side work. Additional methods and structures are provided to give insight into the latency pipeline of an application through the latency markers. `VK_NV_low_latency` provides legacy support for applications that make use of the NVIDIA Reflex SDK whereas new implementations should use the `VK_NV_low_latency2` extension.

## [](#_issues)Issues

1\) How does Low Latency 2 work with applications that utilize device groups?

Low Latency 2 does not support device groups.

## [](#_version_history)Version History

- Revision 2, 2023-11-15 (Charles Hansen)
  
  - Update vkGetLatencyTimingsNV. This is a breaking API change which brings behavior in line with other array querying commands. More background can be found in [https://github.com/KhronosGroup/Vulkan-Docs/issues/2269](https://github.com/KhronosGroup/Vulkan-Docs/issues/2269)
- Revision 1, 2023-09-25 (Charles Hansen)
  
  - Internal revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_NV_low_latency2).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0