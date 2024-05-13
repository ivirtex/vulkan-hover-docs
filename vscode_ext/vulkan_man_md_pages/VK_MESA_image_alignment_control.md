# VK_MESA_image_alignment_control(3) Manual Page

## Name

VK_MESA_image_alignment_control - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

576

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_special_use" class="anchor"></a>Special Use

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#extendingvulkan-compatibility-specialuse"
  target="_blank" rel="noopener">D3D support</a>

## <a href="#_contact" class="anchor"></a>Contact

- Hans-Kristian Arntzen <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_MESA_image_alignment_control%5D%20@HansKristian-Work%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_MESA_image_alignment_control%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>HansKristian-Work</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2024-05-03

**IP Status**  
No known IP claims.

**Contributors**  
- Hans-Kristian Arntzen, Valve

## <a href="#_description" class="anchor"></a>Description

This extension allows applications to request a narrower alignment for
images than an implementation would otherwise require. Some
implementations internally support multiple image layouts in
`VK_IMAGE_TILING_OPTIMAL`, each with different alignment requirements
and performance trade-offs. In some API layering use cases such as
D3D12, it is beneficial to be able to control the alignment, since
certain alignments for placed resources are guaranteed to be supported,
and emulating that expectation requires unnecessary padding of
allocations.

[VkImageAlignmentControlCreateInfoMESA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageAlignmentControlCreateInfoMESA.html)
**can** be chained to [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html),
requesting that the alignment is no more than the provided alignment. If
the requested alignment is not supported for a given
[VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html), a larger alignment **may**
be returned.

While something similar could be achieved with
[`VK_EXT_image_drm_format_modifier`](VK_EXT_image_drm_format_modifier.html)
in theory, this is not the intended way to use that extension. Format
modifiers are generally used for externally shareable images, and would
not be platform portable. It is also a cumbersome API to use just to
lower the alignment.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html):

  - [VkImageAlignmentControlCreateInfoMESA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageAlignmentControlCreateInfoMESA.html)

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceImageAlignmentControlFeaturesMESA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageAlignmentControlFeaturesMESA.html)

- Extending
  [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html):

  - [VkPhysicalDeviceImageAlignmentControlPropertiesMESA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageAlignmentControlPropertiesMESA.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_MESA_IMAGE_ALIGNMENT_CONTROL_EXTENSION_NAME`

- `VK_MESA_IMAGE_ALIGNMENT_CONTROL_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_IMAGE_ALIGNMENT_CONTROL_CREATE_INFO_MESA`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_ALIGNMENT_CONTROL_FEATURES_MESA`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_ALIGNMENT_CONTROL_PROPERTIES_MESA`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2024-04-05 (Hans-Kristian Arntzen)

  - Initial specification

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_MESA_image_alignment_control"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
