# VK\_KHR\_sampler\_ycbcr\_conversion(3) Manual Page

## Name

VK\_KHR\_sampler\_ycbcr\_conversion - device extension



## [](#_registered_extension_number)Registered Extension Number

157

## [](#_revision)Revision

14

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

     [VK\_KHR\_maintenance1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance1.html)  
     and  
     [VK\_KHR\_bind\_memory2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_bind_memory2.html)  
     and  
     [VK\_KHR\_get\_memory\_requirements2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_memory_requirements2.html)  
     and  
     [VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_api_interactions)API Interactions

- Interacts with VK\_EXT\_debug\_report

## [](#_deprecation_state)Deprecation State

- *Promoted* to [Vulkan 1.1](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.1-promotions)

## [](#_contact)Contact

- Andrew Garrard [\[GitHub\]fluppeteer](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_sampler_ycbcr_conversion%5D%20%40fluppeteer%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_sampler_ycbcr_conversion%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

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

## [](#_description)Description

The use of Y′CBCR sampler conversion is an area in 3D graphics not used by most Vulkan developers. It is mainly used for processing inputs from video decoders and cameras. The use of the extension assumes basic knowledge of Y′CBCR concepts.

This extension provides the ability to perform specified color space conversions during texture sampling operations for the Y′CBCR color space natively. It also adds a selection of multi-planar formats, image aspect plane, and the ability to bind memory to the planes of an image collectively or separately.

## [](#_promotion_to_vulkan_1_1)Promotion to Vulkan 1.1

All functionality in this extension is included in core Vulkan 1.1, with the KHR suffix omitted. However, if Vulkan 1.1 is supported and this extension is not, the `samplerYcbcrConversion` capability is optional. The original type, enum, and command names are still available as aliases of the core functionality.

## [](#_promotion_to_vulkan_1_4)Promotion to Vulkan 1.4

If Vulkan 1.4 is supported, support for the `samplerYcbcrConversion` capability is required.

## [](#_new_object_types)New Object Types

- [VkSamplerYcbcrConversionKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrConversionKHR.html)

## [](#_new_commands)New Commands

- [vkCreateSamplerYcbcrConversionKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateSamplerYcbcrConversionKHR.html)
- [vkDestroySamplerYcbcrConversionKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkDestroySamplerYcbcrConversionKHR.html)

## [](#_new_structures)New Structures

- [VkSamplerYcbcrConversionCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrConversionCreateInfoKHR.html)
- Extending [VkBindImageMemoryInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindImageMemoryInfo.html):
  
  - [VkBindImagePlaneMemoryInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindImagePlaneMemoryInfoKHR.html)
- Extending [VkImageFormatProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageFormatProperties2.html):
  
  - [VkSamplerYcbcrConversionImageFormatPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrConversionImageFormatPropertiesKHR.html)
- Extending [VkImageMemoryRequirementsInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageMemoryRequirementsInfo2.html):
  
  - [VkImagePlaneMemoryRequirementsInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImagePlaneMemoryRequirementsInfoKHR.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceSamplerYcbcrConversionFeaturesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceSamplerYcbcrConversionFeaturesKHR.html)
- Extending [VkSamplerCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerCreateInfo.html), [VkImageViewCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewCreateInfo.html):
  
  - [VkSamplerYcbcrConversionInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrConversionInfoKHR.html)

## [](#_new_enums)New Enums

- [VkChromaLocationKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkChromaLocationKHR.html)
- [VkSamplerYcbcrModelConversionKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrModelConversionKHR.html)
- [VkSamplerYcbcrRangeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrRangeKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_SAMPLER_YCBCR_CONVERSION_EXTENSION_NAME`
- `VK_KHR_SAMPLER_YCBCR_CONVERSION_SPEC_VERSION`
- Extending [VkChromaLocation](https://registry.khronos.org/vulkan/specs/latest/man/html/VkChromaLocation.html):
  
  - `VK_CHROMA_LOCATION_COSITED_EVEN_KHR`
  - `VK_CHROMA_LOCATION_MIDPOINT_KHR`
- Extending [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html):
  
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
- Extending [VkFormatFeatureFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatFeatureFlagBits.html):
  
  - `VK_FORMAT_FEATURE_COSITED_CHROMA_SAMPLES_BIT_KHR`
  - `VK_FORMAT_FEATURE_DISJOINT_BIT_KHR`
  - `VK_FORMAT_FEATURE_MIDPOINT_CHROMA_SAMPLES_BIT_KHR`
  - `VK_FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_CHROMA_RECONSTRUCTION_EXPLICIT_BIT_KHR`
  - `VK_FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_CHROMA_RECONSTRUCTION_EXPLICIT_FORCEABLE_BIT_KHR`
  - `VK_FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_LINEAR_FILTER_BIT_KHR`
  - `VK_FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_SEPARATE_RECONSTRUCTION_FILTER_BIT_KHR`
- Extending [VkImageAspectFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageAspectFlagBits.html):
  
  - `VK_IMAGE_ASPECT_PLANE_0_BIT_KHR`
  - `VK_IMAGE_ASPECT_PLANE_1_BIT_KHR`
  - `VK_IMAGE_ASPECT_PLANE_2_BIT_KHR`
- Extending [VkImageCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateFlagBits.html):
  
  - `VK_IMAGE_CREATE_DISJOINT_BIT_KHR`
- Extending [VkObjectType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkObjectType.html):
  
  - `VK_OBJECT_TYPE_SAMPLER_YCBCR_CONVERSION_KHR`
- Extending [VkSamplerYcbcrModelConversion](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrModelConversion.html):
  
  - `VK_SAMPLER_YCBCR_MODEL_CONVERSION_RGB_IDENTITY_KHR`
  - `VK_SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_2020_KHR`
  - `VK_SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_601_KHR`
  - `VK_SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_709_KHR`
  - `VK_SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_IDENTITY_KHR`
- Extending [VkSamplerYcbcrRange](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrRange.html):
  
  - `VK_SAMPLER_YCBCR_RANGE_ITU_FULL_KHR`
  - `VK_SAMPLER_YCBCR_RANGE_ITU_NARROW_KHR`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_BIND_IMAGE_PLANE_MEMORY_INFO_KHR`
  - `VK_STRUCTURE_TYPE_IMAGE_PLANE_MEMORY_REQUIREMENTS_INFO_KHR`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SAMPLER_YCBCR_CONVERSION_FEATURES_KHR`
  - `VK_STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_CREATE_INFO_KHR`
  - `VK_STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_IMAGE_FORMAT_PROPERTIES_KHR`
  - `VK_STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_INFO_KHR`

If [VK\_EXT\_debug\_report](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_debug_report.html) is supported:

- Extending [VkDebugReportObjectTypeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugReportObjectTypeEXT.html):
  
  - `VK_DEBUG_REPORT_OBJECT_TYPE_SAMPLER_YCBCR_CONVERSION_EXT`
  - `VK_DEBUG_REPORT_OBJECT_TYPE_SAMPLER_YCBCR_CONVERSION_KHR_EXT`

## [](#_version_history)Version History

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
  
  - Added cosited option/midpoint requirement for formats, “bypassConversion”
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

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_sampler_ycbcr_conversion)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0