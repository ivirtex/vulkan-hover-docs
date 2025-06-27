# VK_KHR_maintenance1(3) Manual Page

## Name

VK_KHR_maintenance1 - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

70

## <a href="#_revision" class="anchor"></a>Revision

2

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

None

## <a href="#_deprecation_state" class="anchor"></a>Deprecation State

- *Promoted* to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#versions-1.1-promotions"
  target="_blank" rel="noopener">Vulkan 1.1</a>

## <a href="#_contact" class="anchor"></a>Contact

- Piers Daniell <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_maintenance1%5D%20@pdaniell-nv%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_maintenance1%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>pdaniell-nv</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

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

## <a href="#_description" class="anchor"></a>Description

`VK_KHR_maintenance1` adds a collection of minor features that were
intentionally left out or overlooked from the original Vulkan 1.0
release.

The new features are as follows:

- Allow 2D and 2D array image views to be created from 3D images, which
  can then be used as color framebuffer attachments. This allows
  applications to render to slices of a 3D image.

- Support [vkCmdCopyImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyImage.html) between 2D array layers
  and 3D slices. This extension allows copying from layers of a 2D array
  image to slices of a 3D image and vice versa.

- Allow negative height to be specified in the
  [VkViewport](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkViewport.html)::`height` field to perform y-inversion
  of the clip-space to framebuffer-space transform. This allows apps to
  avoid having to use `gl_Position.y = -gl_Position.y` in shaders also
  targeting other APIs.

- Allow implementations to express support for doing just transfers and
  clears of image formats that they otherwise support no other format
  features for. This is done by adding new format feature flags
  `VK_FORMAT_FEATURE_TRANSFER_SRC_BIT_KHR` and
  `VK_FORMAT_FEATURE_TRANSFER_DST_BIT_KHR`.

- Support [vkCmdFillBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdFillBuffer.html) on transfer-only
  queues. Previously [vkCmdFillBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdFillBuffer.html) was defined
  to only work on command buffers allocated from command pools which
  support graphics or compute queues. It is now allowed on queues that
  just support transfer operations.

- Fix the inconsistency of how error conditions are returned between the
  [vkCreateGraphicsPipelines](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateGraphicsPipelines.html) and
  [vkCreateComputePipelines](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateComputePipelines.html) functions
  and the [vkAllocateDescriptorSets](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkAllocateDescriptorSets.html) and
  [vkAllocateCommandBuffers](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkAllocateCommandBuffers.html) functions.

- Add new `VK_ERROR_OUT_OF_POOL_MEMORY_KHR` error so implementations can
  give a more precise reason for
  [vkAllocateDescriptorSets](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkAllocateDescriptorSets.html) failures.

- Add a new command [vkTrimCommandPoolKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkTrimCommandPoolKHR.html)
  which gives the implementation an opportunity to release any unused
  command pool memory back to the system.

## <a href="#_promotion_to_vulkan_1_1" class="anchor"></a>Promotion to Vulkan 1.1

All functionality in this extension is included in core Vulkan 1.1, with
the KHR suffix omitted. The original type, enum and command names are
still available as aliases of the core functionality.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkTrimCommandPoolKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkTrimCommandPoolKHR.html)

## <a href="#_new_bitmasks" class="anchor"></a>New Bitmasks

- [VkCommandPoolTrimFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandPoolTrimFlagsKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_MAINTENANCE1_EXTENSION_NAME`

- `VK_KHR_MAINTENANCE1_SPEC_VERSION`

- `VK_KHR_MAINTENANCE_1_EXTENSION_NAME`

- `VK_KHR_MAINTENANCE_1_SPEC_VERSION`

- Extending [VkFormatFeatureFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatFeatureFlagBits.html):

  - `VK_FORMAT_FEATURE_TRANSFER_DST_BIT_KHR`

  - `VK_FORMAT_FEATURE_TRANSFER_SRC_BIT_KHR`

- Extending [VkImageCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateFlagBits.html):

  - `VK_IMAGE_CREATE_2D_ARRAY_COMPATIBLE_BIT_KHR`

- Extending [VkResult](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResult.html):

  - `VK_ERROR_OUT_OF_POOL_MEMORY_KHR`

## <a href="#_issues" class="anchor"></a>Issues

1.  Are viewports with zero height allowed?

    **RESOLVED**: Yes, although they have low utility.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2016-10-26 (Piers Daniell)

  - Internal revisions

- Revision 2, 2018-03-13 (Jon Leech)

  - Add issue for zero-height viewports

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_maintenance1"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
