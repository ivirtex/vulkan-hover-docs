# VK_KHR_sampler_ycbcr_conversion(3) Manual Page

## Name

VK_KHR_sampler_ycbcr_conversion - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

157

## <a href="#_revision" class="anchor"></a>Revision

14

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

     [VK_KHR_maintenance1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_maintenance1.html)  
     and  
     [VK_KHR_bind_memory2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_bind_memory2.html)  
     and  
    
[VK_KHR_get_memory_requirements2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_memory_requirements2.html)  
     and  
    
[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_api_interactions" class="anchor"></a>API Interactions

- Interacts with VK_EXT_debug_report

## <a href="#_deprecation_state" class="anchor"></a>Deprecation State

- *Promoted* to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#versions-1.1-promotions"
  target="_blank" rel="noopener">Vulkan 1.1</a>

## <a href="#_contact" class="anchor"></a>Contact

- Andrew Garrard <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_sampler_ycbcr_conversion%5D%20@fluppeteer%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_sampler_ycbcr_conversion%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>fluppeteer</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2017-08-11

**IP Status**  
No known IP claims.

**Contributors**  
- Andrew Garrard, Samsung Electronics

- Tobias Hector, Imagination Technologies

- James Jones, NVIDIA

- Daniel Koch, NVIDIA

- Daniel Rakos, AMD

- Romain Guy, Google

- Jesse Hall, Google

- Tom Cooksey, ARM Ltd

- Jeff Leger, Qualcomm Technologies, Inc

- Jan-Harald Fredriksen, ARM Ltd

- Jan Outters, Samsung Electronics

- Alon Or-bach, Samsung Electronics

- Michael Worcester, Imagination Technologies

- Jeff Bolz, NVIDIA

- Tony Zlatinski, NVIDIA

- Matthew Netsch, Qualcomm Technologies, Inc

## <a href="#_description" class="anchor"></a>Description

The use of Y′C<sub>B</sub>C<sub>R</sub> sampler conversion is an area in
3D graphics not used by most Vulkan developers. It is mainly used for
processing inputs from video decoders and cameras. The use of the
extension assumes basic knowledge of Y′C<sub>B</sub>C<sub>R</sub>
concepts.

This extension provides the ability to perform specified color space
conversions during texture sampling operations for the
Y′C<sub>B</sub>C<sub>R</sub> color space natively. It also adds a
selection of multi-planar formats, image aspect plane, and the ability
to bind memory to the planes of an image collectively or separately.

## <a href="#_promotion_to_vulkan_1_1" class="anchor"></a>Promotion to Vulkan 1.1

All functionality in this extension is included in core Vulkan 1.1, with
the KHR suffix omitted. However, if Vulkan 1.1 is supported and this
extension is not, the `samplerYcbcrConversion` capability is optional.
The original type, enum and command names are still available as aliases
of the core functionality.

## <a href="#_new_object_types" class="anchor"></a>New Object Types

- [VkSamplerYcbcrConversionKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrConversionKHR.html)

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkCreateSamplerYcbcrConversionKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateSamplerYcbcrConversionKHR.html)

- [vkDestroySamplerYcbcrConversionKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroySamplerYcbcrConversionKHR.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkSamplerYcbcrConversionCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrConversionCreateInfoKHR.html)

- Extending [VkBindImageMemoryInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindImageMemoryInfo.html):

  - [VkBindImagePlaneMemoryInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindImagePlaneMemoryInfoKHR.html)

- Extending [VkImageFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageFormatProperties2.html):

  - [VkSamplerYcbcrConversionImageFormatPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrConversionImageFormatPropertiesKHR.html)

- Extending
  [VkImageMemoryRequirementsInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageMemoryRequirementsInfo2.html):

  - [VkImagePlaneMemoryRequirementsInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImagePlaneMemoryRequirementsInfoKHR.html)

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceSamplerYcbcrConversionFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSamplerYcbcrConversionFeaturesKHR.html)

- Extending [VkSamplerCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerCreateInfo.html),
  [VkImageViewCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewCreateInfo.html):

  - [VkSamplerYcbcrConversionInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrConversionInfoKHR.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkChromaLocationKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkChromaLocationKHR.html)

- [VkSamplerYcbcrModelConversionKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrModelConversionKHR.html)

- [VkSamplerYcbcrRangeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrRangeKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_SAMPLER_YCBCR_CONVERSION_EXTENSION_NAME`

- `VK_KHR_SAMPLER_YCBCR_CONVERSION_SPEC_VERSION`

- Extending [VkChromaLocation](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkChromaLocation.html):

  - `VK_CHROMA_LOCATION_COSITED_EVEN_KHR`

  - `VK_CHROMA_LOCATION_MIDPOINT_KHR`

- Extending [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html):

  - `VK_FORMAT_B10X6G10X6R10X6G10X6_422_UNORM_4PACK16_KHR`

  - `VK_FORMAT_B12X4G12X4R12X4G12X4_422_UNORM_4PACK16_KHR`

  - `VK_FORMAT_B16G16R16G16_422_UNORM_KHR`

  - `VK_FORMAT_B8G8R8G8_422_UNORM_KHR`

  - `VK_FORMAT_G10X6B10X6G10X6R10X6_422_UNORM_4PACK16_KHR`

  - `VK_FORMAT_G10X6_B10X6R10X6_2PLANE_420_UNORM_3PACK16_KHR`

  - `VK_FORMAT_G10X6_B10X6R10X6_2PLANE_422_UNORM_3PACK16_KHR`

  - `VK_FORMAT_G10X6_B10X6_R10X6_3PLANE_420_UNORM_3PACK16_KHR`

  - `VK_FORMAT_G10X6_B10X6_R10X6_3PLANE_422_UNORM_3PACK16_KHR`

  - `VK_FORMAT_G10X6_B10X6_R10X6_3PLANE_444_UNORM_3PACK16_KHR`

  - `VK_FORMAT_G12X4B12X4G12X4R12X4_422_UNORM_4PACK16_KHR`

  - `VK_FORMAT_G12X4_B12X4R12X4_2PLANE_420_UNORM_3PACK16_KHR`

  - `VK_FORMAT_G12X4_B12X4R12X4_2PLANE_422_UNORM_3PACK16_KHR`

  - `VK_FORMAT_G12X4_B12X4_R12X4_3PLANE_420_UNORM_3PACK16_KHR`

  - `VK_FORMAT_G12X4_B12X4_R12X4_3PLANE_422_UNORM_3PACK16_KHR`

  - `VK_FORMAT_G12X4_B12X4_R12X4_3PLANE_444_UNORM_3PACK16_KHR`

  - `VK_FORMAT_G16B16G16R16_422_UNORM_KHR`

  - `VK_FORMAT_G16_B16R16_2PLANE_420_UNORM_KHR`

  - `VK_FORMAT_G16_B16R16_2PLANE_422_UNORM_KHR`

  - `VK_FORMAT_G16_B16_R16_3PLANE_420_UNORM_KHR`

  - `VK_FORMAT_G16_B16_R16_3PLANE_422_UNORM_KHR`

  - `VK_FORMAT_G16_B16_R16_3PLANE_444_UNORM_KHR`

  - `VK_FORMAT_G8B8G8R8_422_UNORM_KHR`

  - `VK_FORMAT_G8_B8R8_2PLANE_420_UNORM_KHR`

  - `VK_FORMAT_G8_B8R8_2PLANE_422_UNORM_KHR`

  - `VK_FORMAT_G8_B8_R8_3PLANE_420_UNORM_KHR`

  - `VK_FORMAT_G8_B8_R8_3PLANE_422_UNORM_KHR`

  - `VK_FORMAT_G8_B8_R8_3PLANE_444_UNORM_KHR`

  - `VK_FORMAT_R10X6G10X6B10X6A10X6_UNORM_4PACK16_KHR`

  - `VK_FORMAT_R10X6G10X6_UNORM_2PACK16_KHR`

  - `VK_FORMAT_R10X6_UNORM_PACK16_KHR`

  - `VK_FORMAT_R12X4G12X4B12X4A12X4_UNORM_4PACK16_KHR`

  - `VK_FORMAT_R12X4G12X4_UNORM_2PACK16_KHR`

  - `VK_FORMAT_R12X4_UNORM_PACK16_KHR`

- Extending [VkFormatFeatureFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatFeatureFlagBits.html):

  - `VK_FORMAT_FEATURE_COSITED_CHROMA_SAMPLES_BIT_KHR`

  - `VK_FORMAT_FEATURE_DISJOINT_BIT_KHR`

  - `VK_FORMAT_FEATURE_MIDPOINT_CHROMA_SAMPLES_BIT_KHR`

  - `VK_FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_CHROMA_RECONSTRUCTION_EXPLICIT_BIT_KHR`

  - `VK_FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_CHROMA_RECONSTRUCTION_EXPLICIT_FORCEABLE_BIT_KHR`

  - `VK_FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_LINEAR_FILTER_BIT_KHR`

  - `VK_FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_SEPARATE_RECONSTRUCTION_FILTER_BIT_KHR`

- Extending [VkImageAspectFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageAspectFlagBits.html):

  - `VK_IMAGE_ASPECT_PLANE_0_BIT_KHR`

  - `VK_IMAGE_ASPECT_PLANE_1_BIT_KHR`

  - `VK_IMAGE_ASPECT_PLANE_2_BIT_KHR`

- Extending [VkImageCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateFlagBits.html):

  - `VK_IMAGE_CREATE_DISJOINT_BIT_KHR`

- Extending [VkObjectType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkObjectType.html):

  - `VK_OBJECT_TYPE_SAMPLER_YCBCR_CONVERSION_KHR`

- Extending
  [VkSamplerYcbcrModelConversion](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrModelConversion.html):

  - `VK_SAMPLER_YCBCR_MODEL_CONVERSION_RGB_IDENTITY_KHR`

  - `VK_SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_2020_KHR`

  - `VK_SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_601_KHR`

  - `VK_SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_709_KHR`

  - `VK_SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_IDENTITY_KHR`

- Extending [VkSamplerYcbcrRange](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrRange.html):

  - `VK_SAMPLER_YCBCR_RANGE_ITU_FULL_KHR`

  - `VK_SAMPLER_YCBCR_RANGE_ITU_NARROW_KHR`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_BIND_IMAGE_PLANE_MEMORY_INFO_KHR`

  - `VK_STRUCTURE_TYPE_IMAGE_PLANE_MEMORY_REQUIREMENTS_INFO_KHR`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SAMPLER_YCBCR_CONVERSION_FEATURES_KHR`

  - `VK_STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_CREATE_INFO_KHR`

  - `VK_STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_IMAGE_FORMAT_PROPERTIES_KHR`

  - `VK_STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_INFO_KHR`

If [VK_EXT_debug_report](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_debug_report.html) is supported:

- Extending
  [VkDebugReportObjectTypeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugReportObjectTypeEXT.html):

  - `VK_DEBUG_REPORT_OBJECT_TYPE_SAMPLER_YCBCR_CONVERSION_EXT`

  - `VK_DEBUG_REPORT_OBJECT_TYPE_SAMPLER_YCBCR_CONVERSION_KHR_EXT`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2017-01-24 (Andrew Garrard)

  - Initial draft

- Revision 2, 2017-01-25 (Andrew Garrard)

  - After initial feedback

- Revision 3, 2017-01-27 (Andrew Garrard)

  - Higher bit depth formats, renaming, swizzle

- Revision 4, 2017-02-22 (Andrew Garrard)

  - Added query function, formats as RGB, clarifications

- Revision 5, 2017-04-?? (Andrew Garrard)

  - Simplified query and removed output conversions

- Revision 6, 2017-04-24 (Andrew Garrard)

  - Tidying, incorporated new image query, restored transfer functions

- Revision 7, 2017-04-25 (Andrew Garrard)

  - Added cosited option/midpoint requirement for formats,
    “bypassConversion”

- Revision 8, 2017-04-25 (Andrew Garrard)

  - Simplified further

- Revision 9, 2017-04-27 (Andrew Garrard)

  - Disjoint no more

- Revision 10, 2017-04-28 (Andrew Garrard)

  - Restored disjoint

- Revision 11, 2017-04-29 (Andrew Garrard)

  - Now Ycbcr conversion, and KHR

- Revision 12, 2017-06-06 (Andrew Garrard)

  - Added conversion to image view creation

- Revision 13, 2017-07-13 (Andrew Garrard)

  - Allowed cosited-only chroma samples for formats

- Revision 14, 2017-08-11 (Andrew Garrard)

  - Reflected quantization changes in BT.2100-1

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_sampler_ycbcr_conversion"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
