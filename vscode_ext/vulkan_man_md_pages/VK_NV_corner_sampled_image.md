# VK_NV_corner_sampled_image(3) Manual Page

## Name

VK_NV_corner_sampled_image - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

51

## <a href="#_revision" class="anchor"></a>Revision

2

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_contact" class="anchor"></a>Contact

- Daniel Koch <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_corner_sampled_image%5D%20@dgkoch%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_corner_sampled_image%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>dgkoch</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2018-08-13

**Contributors**  
- Jeff Bolz, NVIDIA

- Pat Brown, NVIDIA

- Chris Lentini, NVIDIA

## <a href="#_description" class="anchor"></a>Description

This extension adds support for a new image organization, which this
extension refers to as “corner-sampled” images. A corner-sampled image
differs from a conventional image in the following ways:

- Texels are centered on integer coordinates. See <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#textures-unnormalized-to-integer"
  target="_blank" rel="noopener">Unnormalized Texel Coordinate
  Operations</a>

- Normalized coordinates are scaled using coord × (dim - 1) rather than
  coord × dim, where dim is the size of one dimension of the image. See
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#textures-normalized-to-unnormalized"
  target="_blank" rel="noopener">normalized texel coordinate transform</a>.

- Partial derivatives are scaled using coord × (dim - 1) rather than
  coord × dim. See <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#textures-scale-factor"
  target="_blank" rel="noopener">Scale Factor Operation</a>.

- Calculation of the next higher LOD size goes according to ⌈dim / 2⌉
  rather than ⌊dim / 2⌋. See <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-image-mip-level-sizing"
  target="_blank" rel="noopener">Image Mip Level Sizing</a>.

- The minimum level size is 2x2 for 2D images and 2x2x2 for 3D images.
  See <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-image-mip-level-sizing"
  target="_blank" rel="noopener">Image Mip Level Sizing</a>.

This image organization is designed to facilitate a system like Ptex
with separate textures for each face of a subdivision or polygon mesh.
Placing sample locations at pixel corners allows applications to
maintain continuity between adjacent patches by duplicating values along
shared edges. Additionally, using the modified mipmapping logic along
with texture dimensions of the form 2<sup>n</sup>+1 allows continuity
across shared edges even if the adjacent patches use different LOD
values.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceCornerSampledImageFeaturesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceCornerSampledImageFeaturesNV.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_NV_CORNER_SAMPLED_IMAGE_EXTENSION_NAME`

- `VK_NV_CORNER_SAMPLED_IMAGE_SPEC_VERSION`

- Extending [VkImageCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateFlagBits.html):

  - `VK_IMAGE_CREATE_CORNER_SAMPLED_BIT_NV`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_CORNER_SAMPLED_IMAGE_FEATURES_NV`

## <a href="#_issues" class="anchor"></a>Issues

1.  What should this extension be named?

    **DISCUSSION**: While naming this extension, we chose the most
    distinctive aspect of the image organization and referred to such
    images as “corner-sampled images”. As a result, we decided to name
    the extension NV_corner_sampled_image.

2.  Do we need a format feature flag so formats can advertise if they
    support corner-sampling?

    **DISCUSSION**: Currently NVIDIA supports this for all 2D and 3D
    formats, but not for cube maps or depth-stencil formats. A format
    feature might be useful if other vendors would only support this on
    some formats.

3.  Do integer texel coordinates have a different range for
    corner-sampled images?

    **RESOLVED**: No, these are unchanged.

4.  Do unnormalized sampler coordinates work with corner-sampled images?
    Are there any functional differences?

    **RESOLVED**: Yes. Unnormalized coordinates are treated as already
    scaled for corner-sample usage.

5.  Should we have a diagram in the “Image Operations” chapter
    demonstrating different texel sampling locations?

    **UNRESOLVED**: Probably, but later.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2018-08-14 (Daniel Koch)

  - Internal revisions

- Revision 2, 2018-08-14 (Daniel Koch)

  - ???

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_NV_corner_sampled_image"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
