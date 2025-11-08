# VK\_VALVE\_video\_encode\_rgb\_conversion(3) Manual Page

## Name

VK\_VALVE\_video\_encode\_rgb\_conversion - device extension



## [](#_registered_extension_number)Registered Extension Number

391

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_video\_encode\_queue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_encode_queue.html)  
and  
     [VK\_KHR\_sampler\_ycbcr\_conversion](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_sampler_ycbcr_conversion.html)  
     or  
     [Vulkan Version 1.1](#versions-1.1)

## [](#_contact)Contact

- Autumn Ashton [\[GitHub\]misyltoad](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_VALVE_video_encode_rgb_conversion%5D%20%40misyltoad%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_VALVE_video_encode_rgb_conversion%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2025-08-22

**IP Status**

No known IP claims.

**Contributors**

- Autumn Ashton, Valve
- Daniel Rakos, RasterGrid

## [](#_description)Description

This extension builds upon the `VK_KHR_video_encode_queue` extension by enabling the application to pass in RGB/RGBA images in video encode operations.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceVideoEncodeRgbConversionFeaturesVALVE](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceVideoEncodeRgbConversionFeaturesVALVE.html)
- Extending [VkVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoCapabilitiesKHR.html):
  
  - [VkVideoEncodeRgbConversionCapabilitiesVALVE](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeRgbConversionCapabilitiesVALVE.html)
- Extending [VkVideoProfileInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoProfileInfoKHR.html):
  
  - [VkVideoEncodeProfileRgbConversionInfoVALVE](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeProfileRgbConversionInfoVALVE.html)
- Extending [VkVideoSessionCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionCreateInfoKHR.html):
  
  - [VkVideoEncodeSessionRgbConversionCreateInfoVALVE](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeSessionRgbConversionCreateInfoVALVE.html)

## [](#_new_enums)New Enums

- [VkVideoEncodeRgbChromaOffsetFlagBitsVALVE](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeRgbChromaOffsetFlagBitsVALVE.html)
- [VkVideoEncodeRgbModelConversionFlagBitsVALVE](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeRgbModelConversionFlagBitsVALVE.html)
- [VkVideoEncodeRgbRangeCompressionFlagBitsVALVE](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeRgbRangeCompressionFlagBitsVALVE.html)

## [](#_new_bitmasks)New Bitmasks

- [VkVideoEncodeRgbChromaOffsetFlagsVALVE](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeRgbChromaOffsetFlagsVALVE.html)
- [VkVideoEncodeRgbModelConversionFlagsVALVE](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeRgbModelConversionFlagsVALVE.html)
- [VkVideoEncodeRgbRangeCompressionFlagsVALVE](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoEncodeRgbRangeCompressionFlagsVALVE.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_VALVE_VIDEO_ENCODE_RGB_CONVERSION_EXTENSION_NAME`
- `VK_VALVE_VIDEO_ENCODE_RGB_CONVERSION_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VIDEO_ENCODE_RGB_CONVERSION_FEATURES_VALVE`
  - `VK_STRUCTURE_TYPE_VIDEO_ENCODE_PROFILE_RGB_CONVERSION_INFO_VALVE`
  - `VK_STRUCTURE_TYPE_VIDEO_ENCODE_RGB_CONVERSION_CAPABILITIES_VALVE`
  - `VK_STRUCTURE_TYPE_VIDEO_ENCODE_SESSION_RGB_CONVERSION_CREATE_INFO_VALVE`

## [](#_version_history)Version History

- Revision 1, 2025-08-22 (Autumn Ashton)
  
  - Initial version

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_VALVE_video_encode_rgb_conversion).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0