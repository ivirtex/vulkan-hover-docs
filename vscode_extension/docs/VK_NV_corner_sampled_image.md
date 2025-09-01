# VK\_NV\_corner\_sampled\_image(3) Manual Page

## Name

VK\_NV\_corner\_sampled\_image - device extension



## [](#_registered_extension_number)Registered Extension Number

51

## [](#_revision)Revision

2

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_contact)Contact

- Daniel Koch [\[GitHub\]dgkoch](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_corner_sampled_image%5D%20%40dgkoch%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_corner_sampled_image%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2018-08-13

**Contributors**

- Jeff Bolz, NVIDIA
- Pat Brown, NVIDIA
- Chris Lentini, NVIDIA

## [](#_description)Description

This extension adds support for a new image organization, which this extension refers to as “corner-sampled” images. A corner-sampled image differs from a conventional image in the following ways:

- Texels are centered on integer coordinates. See [Unnormalized Texel Coordinate Operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#textures-unnormalized-to-integer)
- Normalized coordinates are scaled using coord × (dim - 1) rather than coord × dim, where dim is the size of one dimension of the image. See [normalized texel coordinate transform](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#textures-normalized-to-unnormalized).
- Partial derivatives are scaled using coord × (dim - 1) rather than coord × dim. See [Scale Factor Operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#textures-scale-factor).
- Calculation of the next higher LOD size goes according to ⌈dim / 2⌉ rather than ⌊dim / 2⌋. See [Image Mip Level Sizing](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#resources-image-mip-level-sizing).
- The minimum level size is 2x2 for 2D images and 2x2x2 for 3D images. See [Image Mip Level Sizing](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#resources-image-mip-level-sizing).

This image organization is designed to facilitate a system like Ptex with separate textures for each face of a subdivision or polygon mesh. Placing sample locations at pixel corners allows applications to maintain continuity between adjacent patches by duplicating values along shared edges. Additionally, using the modified mipmapping logic along with texture dimensions of the form 2n+1 allows continuity across shared edges even if the adjacent patches use different LOD values.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceCornerSampledImageFeaturesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceCornerSampledImageFeaturesNV.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_NV_CORNER_SAMPLED_IMAGE_EXTENSION_NAME`
- `VK_NV_CORNER_SAMPLED_IMAGE_SPEC_VERSION`
- Extending [VkImageCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateFlagBits.html):
  
  - `VK_IMAGE_CREATE_CORNER_SAMPLED_BIT_NV`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_CORNER_SAMPLED_IMAGE_FEATURES_NV`

## [](#_issues)Issues

1. What should this extension be named?
   
   **DISCUSSION**: While naming this extension, we chose the most distinctive aspect of the image organization and referred to such images as “corner-sampled images”. As a result, we decided to name the extension NV\_corner\_sampled\_image.
2. Do we need a format feature flag so formats can advertise if they support corner-sampling?
   
   **DISCUSSION**: Currently NVIDIA supports this for all 2D and 3D formats, but not for cube maps or depth-stencil formats. A format feature might be useful if other vendors would only support this on some formats.
3. Do integer texel coordinates have a different range for corner-sampled images?
   
   **RESOLVED**: No, these are unchanged.
4. Do unnormalized sampler coordinates work with corner-sampled images? Are there any functional differences?
   
   **RESOLVED**: Yes. Unnormalized coordinates are treated as already scaled for corner-sample usage.
5. Should we have a diagram in the “Image Operations” chapter demonstrating different texel sampling locations?
   
   **UNRESOLVED**: Probably, but later.

## [](#_version_history)Version History

- Revision 1, 2018-08-14 (Daniel Koch)
  
  - Internal revisions
- Revision 2, 2018-08-14 (Daniel Koch)
  
  - ???

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_NV_corner_sampled_image).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0