# VK\_KHR\_maintenance1(3) Manual Page

## Name

VK\_KHR\_maintenance1 - device extension



## [](#_registered_extension_number)Registered Extension Number

70

## [](#_revision)Revision

2

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

None

## [](#_deprecation_state)Deprecation State

- *Promoted* to [Vulkan 1.1](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.1-promotions)

## [](#_contact)Contact

- Piers Daniell [\[GitHub\]pdaniell-nv](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_maintenance1%5D%20%40pdaniell-nv%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_maintenance1%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2018-03-13

**Contributors**

- Dan Ginsburg, Valve
- Daniel Koch, NVIDIA
- Daniel Rakos, AMD
- Jan-Harald Fredriksen, ARM
- Faith Ekstrand, Intel
- Jeff Bolz, NVIDIA
- Jesse Hall, Google
- John Kessenich, Google
- Michael Worcester, Imagination Technologies
- Neil Henning, Codeplay Software Ltd.
- Piers Daniell, NVIDIA
- Slawomir Grajewski, Intel
- Tobias Hector, Imagination Technologies
- Tom Olson, ARM

## [](#_description)Description

`VK_KHR_maintenance1` adds a collection of minor features that were intentionally left out or overlooked from the original Vulkan 1.0 release.

The new features are as follows:

- Allow 2D and 2D array image views to be created from 3D images, which can then be used as color framebuffer attachments. This allows applications to render to slices of a 3D image.
- Support [vkCmdCopyImage](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyImage.html) between 2D array layers and 3D slices. This extension allows copying from layers of a 2D array image to slices of a 3D image and vice versa.
- Allow negative height to be specified in the [VkViewport](https://registry.khronos.org/vulkan/specs/latest/man/html/VkViewport.html)::`height` field to perform y-inversion of the clip-space to framebuffer-space transform. This allows apps to avoid having to use `gl_Position.y = -gl_Position.y` in shaders also targeting other APIs.
- Allow implementations to express support for doing just transfers and clears of image formats that they otherwise support no other format features for. This is done by adding new format feature flags `VK_FORMAT_FEATURE_TRANSFER_SRC_BIT_KHR` and `VK_FORMAT_FEATURE_TRANSFER_DST_BIT_KHR`.
- Support [vkCmdFillBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdFillBuffer.html) on transfer-only queues. Previously [vkCmdFillBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdFillBuffer.html) was defined to only work on command buffers allocated from command pools which support graphics or compute queues. It is now allowed on queues that just support transfer operations.
- Fix the inconsistency of how error conditions are returned between the [vkCreateGraphicsPipelines](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateGraphicsPipelines.html) and [vkCreateComputePipelines](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateComputePipelines.html) functions and the [vkAllocateDescriptorSets](https://registry.khronos.org/vulkan/specs/latest/man/html/vkAllocateDescriptorSets.html) and [vkAllocateCommandBuffers](https://registry.khronos.org/vulkan/specs/latest/man/html/vkAllocateCommandBuffers.html) functions.
- Add new `VK_ERROR_OUT_OF_POOL_MEMORY_KHR` error so implementations can give a more precise reason for [vkAllocateDescriptorSets](https://registry.khronos.org/vulkan/specs/latest/man/html/vkAllocateDescriptorSets.html) failures.
- Add a new command [vkTrimCommandPoolKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkTrimCommandPoolKHR.html) which gives the implementation an opportunity to release any unused command pool memory back to the system.

## [](#_promotion_to_vulkan_1_1)Promotion to Vulkan 1.1

All functionality in this extension is included in core Vulkan 1.1, with the KHR suffix omitted. The original type, enum, and command names are still available as aliases of the core functionality.

## [](#_new_commands)New Commands

- [vkTrimCommandPoolKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkTrimCommandPoolKHR.html)

## [](#_new_bitmasks)New Bitmasks

- [VkCommandPoolTrimFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandPoolTrimFlagsKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_MAINTENANCE1_EXTENSION_NAME`
- `VK_KHR_MAINTENANCE1_SPEC_VERSION`
- `VK_KHR_MAINTENANCE_1_EXTENSION_NAME`
- `VK_KHR_MAINTENANCE_1_SPEC_VERSION`
- Extending [VkFormatFeatureFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatFeatureFlagBits.html):
  
  - `VK_FORMAT_FEATURE_TRANSFER_DST_BIT_KHR`
  - `VK_FORMAT_FEATURE_TRANSFER_SRC_BIT_KHR`
- Extending [VkImageCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateFlagBits.html):
  
  - `VK_IMAGE_CREATE_2D_ARRAY_COMPATIBLE_BIT_KHR`
- Extending [VkResult](https://registry.khronos.org/vulkan/specs/latest/man/html/VkResult.html):
  
  - `VK_ERROR_OUT_OF_POOL_MEMORY_KHR`

## [](#_issues)Issues

1. Are viewports with zero height allowed?
   
   **RESOLVED**: Yes, although they have low utility.

## [](#_version_history)Version History

- Revision 1, 2016-10-26 (Piers Daniell)
  
  - Internal revisions
- Revision 2, 2018-03-13 (Jon Leech)
  
  - Add issue for zero-height viewports

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_maintenance1)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0