# VK_KHR_maintenance2(3) Manual Page

## Name

VK_KHR_maintenance2 - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

118

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

None

## <a href="#_deprecation_state" class="anchor"></a>Deprecation State

- *Promoted* to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#versions-1.1-promotions"
  target="_blank" rel="noopener">Vulkan 1.1</a>

## <a href="#_contact" class="anchor"></a>Contact

- Michael Worcester <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_maintenance2%5D%20@michaelworcester%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_maintenance2%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>michaelworcester</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2017-09-05

**Contributors**  
- Michael Worcester, Imagination Technologies

- Stuart Smith, Imagination Technologies

- Jeff Bolz, NVIDIA

- Daniel Koch, NVIDIA

- Jan-Harald Fredriksen, ARM

- Daniel Rakos, AMD

- Neil Henning, Codeplay

- Piers Daniell, NVIDIA

## <a href="#_description" class="anchor"></a>Description

`VK_KHR_maintenance2` adds a collection of minor features that were
intentionally left out or overlooked from the original Vulkan 1.0
release.

The new features are as follows:

- Allow the application to specify which aspect of an input attachment
  might be read for a given subpass.

- Allow implementations to express the clipping behavior of points.

- Allow creating images with usage flags that may not be supported for
  the base image’s format, but are supported for image views of the
  image that have a different but compatible format.

- Allow creating uncompressed image views of compressed images.

- Allow the application to select between an upper-left and lower-left
  origin for the tessellation domain space.

- Adds two new image layouts for depth stencil images to allow either
  the depth or stencil aspect to be read-only while the other aspect is
  writable.

## <a href="#_input_attachment_specification" class="anchor"></a>Input Attachment Specification

Input attachment specification allows an application to specify which
aspect of a multi-aspect image (e.g. a depth/stencil format) will be
accessed via a `subpassLoad` operation.

On some implementations there **may** be a performance penalty if the
implementation does not know (at
[vkCreateRenderPass](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateRenderPass.html) time) which aspect(s) of
multi-aspect images **can** be accessed as input attachments.

## <a href="#_promotion_to_vulkan_1_1" class="anchor"></a>Promotion to Vulkan 1.1

All functionality in this extension is included in core Vulkan 1.1, with
the KHR suffix omitted. The original type, enum and command names are
still available as aliases of the core functionality.

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkInputAttachmentAspectReferenceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInputAttachmentAspectReferenceKHR.html)

- Extending [VkImageViewCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewCreateInfo.html):

  - [VkImageViewUsageCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewUsageCreateInfoKHR.html)

- Extending
  [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html):

  - [VkPhysicalDevicePointClippingPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevicePointClippingPropertiesKHR.html)

- Extending
  [VkPipelineTessellationStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineTessellationStateCreateInfo.html):

  - [VkPipelineTessellationDomainOriginStateCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineTessellationDomainOriginStateCreateInfoKHR.html)

- Extending [VkRenderPassCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassCreateInfo.html):

  - [VkRenderPassInputAttachmentAspectCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassInputAttachmentAspectCreateInfoKHR.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkPointClippingBehaviorKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPointClippingBehaviorKHR.html)

- [VkTessellationDomainOriginKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkTessellationDomainOriginKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_MAINTENANCE2_EXTENSION_NAME`

- `VK_KHR_MAINTENANCE2_SPEC_VERSION`

- `VK_KHR_MAINTENANCE_2_EXTENSION_NAME`

- `VK_KHR_MAINTENANCE_2_SPEC_VERSION`

- Extending [VkImageCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateFlagBits.html):

  - `VK_IMAGE_CREATE_BLOCK_TEXEL_VIEW_COMPATIBLE_BIT_KHR`

  - `VK_IMAGE_CREATE_EXTENDED_USAGE_BIT_KHR`

- Extending [VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html):

  - `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_STENCIL_READ_ONLY_OPTIMAL_KHR`

  - `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_STENCIL_ATTACHMENT_OPTIMAL_KHR`

- Extending [VkPointClippingBehavior](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPointClippingBehavior.html):

  - `VK_POINT_CLIPPING_BEHAVIOR_ALL_CLIP_PLANES_KHR`

  - `VK_POINT_CLIPPING_BEHAVIOR_USER_CLIP_PLANES_ONLY_KHR`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_IMAGE_VIEW_USAGE_CREATE_INFO_KHR`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_POINT_CLIPPING_PROPERTIES_KHR`

  - `VK_STRUCTURE_TYPE_PIPELINE_TESSELLATION_DOMAIN_ORIGIN_STATE_CREATE_INFO_KHR`

  - `VK_STRUCTURE_TYPE_RENDER_PASS_INPUT_ATTACHMENT_ASPECT_CREATE_INFO_KHR`

- Extending
  [VkTessellationDomainOrigin](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkTessellationDomainOrigin.html):

  - `VK_TESSELLATION_DOMAIN_ORIGIN_LOWER_LEFT_KHR`

  - `VK_TESSELLATION_DOMAIN_ORIGIN_UPPER_LEFT_KHR`

## <a href="#_input_attachment_specification_example" class="anchor"></a>Input Attachment Specification Example

Consider the case where a render pass has two subpasses and two
attachments.

Attachment 0 has the format `VK_FORMAT_D24_UNORM_S8_UINT`, attachment 1
has some color format.

Subpass 0 writes to attachment 0, subpass 1 reads only the depth
information from attachment 0 (using inputAttachmentRead) and writes to
attachment 1.

``` c
    VkInputAttachmentAspectReferenceKHR references[] = {
        {
            .subpass = 1,
            .inputAttachmentIndex = 0,
            .aspectMask = VK_IMAGE_ASPECT_DEPTH_BIT
        }
    };

    VkRenderPassInputAttachmentAspectCreateInfoKHR specifyAspects = {
        .sType = VK_STRUCTURE_TYPE_RENDER_PASS_INPUT_ATTACHMENT_ASPECT_CREATE_INFO_KHR,
        .pNext = NULL,
        .aspectReferenceCount = 1,
        .pAspectReferences = references
    };


    VkRenderPassCreateInfo createInfo = {
        ...
        .pNext = &specifyAspects,
        ...
    };

    vkCreateRenderPass(...);
```

## <a href="#_issues" class="anchor"></a>Issues

1\) What is the default tessellation domain origin?

**RESOLVED**: Vulkan 1.0 originally inadvertently documented a
lower-left origin, but the conformance tests and all implementations
implemented an upper-left origin. This extension adds a control to
select between lower-left (for compatibility with OpenGL) and
upper-left, and we retroactively fix unextended Vulkan to have a default
of an upper-left origin.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2017-04-28

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_maintenance2"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
