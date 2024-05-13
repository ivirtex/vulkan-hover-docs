# VK_KHR_separate_depth_stencil_layouts(3) Manual Page

## Name

VK_KHR_separate_depth_stencil_layouts - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

242

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

        
[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
         or  
         [Version 1.1](#versions-1.1)  
     and  
     [VK_KHR_create_renderpass2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_create_renderpass2.html)  
or  
[Version 1.2](#versions-1.2)  

## <a href="#_deprecation_state" class="anchor"></a>Deprecation State

- *Promoted* to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#versions-1.2-promotions"
  target="_blank" rel="noopener">Vulkan 1.2</a>

## <a href="#_contact" class="anchor"></a>Contact

- Piers Daniell <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_separate_depth_stencil_layouts%5D%20@pdaniell-nv%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_separate_depth_stencil_layouts%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>pdaniell-nv</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2019-06-25

**Contributors**  
- Daniel Koch, NVIDIA

- Jeff Bolz, NVIDIA

- Jesse Barker, Unity

- Tobias Hector, AMD

## <a href="#_description" class="anchor"></a>Description

This extension allows image memory barriers for depth/stencil images to
have just one of the `VK_IMAGE_ASPECT_DEPTH_BIT` or
`VK_IMAGE_ASPECT_STENCIL_BIT` aspect bits set, rather than require both.
This allows their layouts to be set independently. To support
depth/stencil images with different layouts for the depth and stencil
aspects, the depth/stencil attachment interface has been updated to
support a separate layout for stencil.

## <a href="#_promotion_to_vulkan_1_2" class="anchor"></a>Promotion to Vulkan 1.2

All functionality in this extension is included in core Vulkan 1.2, with
the KHR suffix omitted. The original type, enum and command names are
still available as aliases of the core functionality.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkAttachmentDescription2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentDescription2.html):

  - [VkAttachmentDescriptionStencilLayoutKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentDescriptionStencilLayoutKHR.html)

- Extending [VkAttachmentReference2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentReference2.html):

  - [VkAttachmentReferenceStencilLayoutKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentReferenceStencilLayoutKHR.html)

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceSeparateDepthStencilLayoutsFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSeparateDepthStencilLayoutsFeaturesKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_SEPARATE_DEPTH_STENCIL_LAYOUTS_EXTENSION_NAME`

- `VK_KHR_SEPARATE_DEPTH_STENCIL_LAYOUTS_SPEC_VERSION`

- Extending [VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html):

  - `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL_KHR`

  - `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL_KHR`

  - `VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL_KHR`

  - `VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL_KHR`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_ATTACHMENT_DESCRIPTION_STENCIL_LAYOUT_KHR`

  - `VK_STRUCTURE_TYPE_ATTACHMENT_REFERENCE_STENCIL_LAYOUT_KHR`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SEPARATE_DEPTH_STENCIL_LAYOUTS_FEATURES_KHR`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2019-06-25 (Piers Daniell)

  - Internal revisions

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_separate_depth_stencil_layouts"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
