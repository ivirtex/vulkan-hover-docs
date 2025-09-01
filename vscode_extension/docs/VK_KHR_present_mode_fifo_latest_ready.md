# VK\_KHR\_present\_mode\_fifo\_latest\_ready(3) Manual Page

## Name

VK\_KHR\_present\_mode\_fifo\_latest\_ready - device extension



## [](#_registered_extension_number)Registered Extension Number

622

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_swapchain](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_swapchain.html)

## [](#_contact)Contact

- Lionel Duc [\[GitHub\]nvlduc](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_present_mode_fifo_latest_ready%5D%20%40nvlduc%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_present_mode_fifo_latest_ready%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_KHR\_present\_mode\_fifo\_latest\_ready](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_KHR_present_mode_fifo_latest_ready.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2025-03-18

**IP Status**

No known IP claims.

**Contributors**

- James Jones, NVIDIA
- Lionel Duc, NVIDIA
- Lina Versace, Google

## [](#_description)Description

This extension is based on `VK_EXT_present_mode_fifo_latest_ready` and provides equivalent functionality.

This extension adds a new present mode, `VK_PRESENT_MODE_FIFO_LATEST_READY_KHR`.

This tear-free present mode behaves much like `VK_PRESENT_MODE_FIFO_KHR`, except that each vertical blanking period dequeues consecutive present requests until the latest ready is found to update the current image.

While this seems similar in concept to `VK_PRESENT_MODE_MAILBOX_KHR`, the fundamental difference is that the processing of the present requests is done during vblank. From the application perspective, this means for example, that in a flip-based model, a single vblank **may** cause multiple swapchain images to be released at once, while `VK_PRESENT_MODE_MAILBOX_KHR` **may** continuously be releasing images as new requests become ready.

This additional present mode is useful when using a time-based present API.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDevicePresentModeFifoLatestReadyFeaturesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevicePresentModeFifoLatestReadyFeaturesKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_PRESENT_MODE_FIFO_LATEST_READY_EXTENSION_NAME`
- `VK_KHR_PRESENT_MODE_FIFO_LATEST_READY_SPEC_VERSION`
- Extending [VkPresentModeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentModeKHR.html):
  
  - `VK_PRESENT_MODE_FIFO_LATEST_READY_KHR`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PRESENT_MODE_FIFO_LATEST_READY_FEATURES_KHR`

## [](#_version_history)Version History

- Revision 1, 2025-03-18 (Lina Versace)
  
  - Internal revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_present_mode_fifo_latest_ready).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0