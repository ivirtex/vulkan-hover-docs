# VK\_KHR\_present\_wait2(3) Manual Page

## Name

VK\_KHR\_present\_wait2 - device extension



## [](#_registered_extension_number)Registered Extension Number

481

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_surface\_capabilities2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_surface_capabilities2.html)  
and  
[VK\_KHR\_surface](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_surface.html)  
and  
[VK\_KHR\_swapchain](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_swapchain.html)  
and  
[VK\_KHR\_present\_id2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_present_id2.html)

## [](#_contact)Contact

- Daniel Stone

## [](#_extension_proposal)Extension Proposal

[VK\_KHR\_present\_wait2](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_KHR_present_wait2.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2025-05-30

**IP Status**

No known IP claims.

**Contributors**

- Hans-Kristian Arntzen, Valve
- James Jones, NVIDIA
- Daniel Stone, Collabora
- Derek Foreman, Collabora
- *contributors to \`[VK\_KHR\_present\_wait](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_present_wait.html)\`*

## [](#_description)Description

This device extension allows an application that uses the `VK_KHR_swapchain` extension to wait for present operations to complete. An application can use this to monitor and control the pacing of the application by managing the number of outstanding images yet to be presented.

## [](#_new_commands)New Commands

- [vkWaitForPresent2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkWaitForPresent2KHR.html)

## [](#_new_structures)New Structures

- [VkPresentWait2InfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentWait2InfoKHR.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDevicePresentWait2FeaturesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevicePresentWait2FeaturesKHR.html)
- Extending [VkSurfaceCapabilities2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceCapabilities2KHR.html):
  
  - [VkSurfaceCapabilitiesPresentWait2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceCapabilitiesPresentWait2KHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_PRESENT_WAIT_2_EXTENSION_NAME`
- `VK_KHR_PRESENT_WAIT_2_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PRESENT_WAIT_2_FEATURES_KHR`
  - `VK_STRUCTURE_TYPE_PRESENT_WAIT_2_INFO_KHR`
  - `VK_STRUCTURE_TYPE_SURFACE_CAPABILITIES_PRESENT_WAIT_2_KHR`
- Extending [VkSwapchainCreateFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainCreateFlagBitsKHR.html):
  
  - `VK_SWAPCHAIN_CREATE_PRESENT_WAIT_2_BIT_KHR`

## [](#_issues)Issues

1\) When does the wait finish?

**RESOLVED**. The wait request will complete when the timeout expires, or after the corresponding presentation request has either taken effect within the presentation engine or has been replaced without presentation. Additionally, a wait may complete immediately if the swapchain becomes out of date.

In circumstances outside the application’s control, this wait may be particularly long. For example, a user session may have the display locked and switched off for multiple days. During this time, the latest image presented through the WSI will never be presented to the user (because nothing is being presented), or replaced (because nothing newer has been queued by the application). Each operating system may have a separate mechanism to inform the application of states such as these, however it is out of scope of the Vulkan WSI.

There is no requirement for any precise timing relationship between the presentation of the image to the user and the end of the wait.

This extension is not intended for time-to-light estimation, which is better solved by a separate extension dedicated to present-timing feedback for audio/visual/input synchronization.

2\) Should this use fences or other existing synchronization mechanism?

**RESOLVED**. VkFence is a legacy primitive. Building a new API around a legacy primitive is undesirable.

Other existing synchronization mechanisms may lack a platform-provided framework for sharing synchronization objects between display and render drivers.

For these reasons, this extension will provide a separate synchronization API.

3\) Should this extension share present identification with other extensions?

**RESOLVED**. Yes. A new extension, `VK_KHR_present_id2`, should be created to provide a shared structure for presentation identifiers.

4\) What happens when presentations complete out of order with respect to calls to vkQueuePresent? This could happen if the semaphores for the presentations were ready out of order.

**OPTION A**: Require that when a PresentId is set that the driver ensure that images are always presented in the order of calls to vkQueuePresent.

**OPTION B**: Finish both waits when the earliest present completes. This will complete the later present wait earlier than the actual presentation. This should be the easiest to implement as the driver need only track the largest present ID completed. This is also the 'natural' consequence of interpreting the existing vkWaitForPresentKHR specification.

**OPTION C**: Finish both waits when both have completed. This will complete the earlier presentation later than the actual presentation time. This is allowed by the current specification as there is no precise timing requirement for when the presentId value is updated. This requires slightly more complexity in the driver as it will need to track all outstanding presentId values.

**OPTION D**: The order of completion between outstanding `vkWaitForPresent2KHR` calls is always undefined. However, a `VK_SUCCESS` return value in [VkPresentWait2InfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentWait2InfoKHR.html)::`presentId` implies that future calls to `vkWaitForPresent2KHR` where [VkPresentWait2InfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentWait2InfoKHR.html)::`presentId` is less than or equal to N will complete immediately.

**RESOLVED**. **OPTION D**: This option ensures implementations do not need to create complex internal queues to generate signals in the right order.

5\) Should this extension deviate from `VK_KHR_present_wait` and require the presentation engine to provide the presentId values?

**RESOLVED**. No. This extension is intended to be a bugfix of `VK_KHR_present_wait`, and existing low-latency apis require an application provided id. At least on some platforms, a mapping mechanism would be required to translate between presentation engine and application ids. This exceeds the intended scope of this extension.

When needed in the future, we can introduce an independent presentation engine driven id and a mechanism for mapping presentation engine ids to application provided ids.

## [](#_examples)Examples

## [](#_version_history)Version History

- Revision 1, 2022-10-05 (Daniel Stone)
  
  - Repurposed from VK\_KHR\_present\_wait to be based on surface capabilities
  - Reworded wait finish section to avoid time-to-light

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_present_wait2)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0