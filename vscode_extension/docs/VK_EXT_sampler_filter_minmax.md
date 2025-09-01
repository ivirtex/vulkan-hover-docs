# VK\_EXT\_sampler\_filter\_minmax(3) Manual Page

## Name

VK\_EXT\_sampler\_filter\_minmax - device extension



## [](#_registered_extension_number)Registered Extension Number

131

## [](#_revision)Revision

2

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_deprecation_state)Deprecation State

- *Promoted* to [Vulkan 1.2](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.2-promotions)

## [](#_contact)Contact

- Jeff Bolz [\[GitHub\]jeffbolznv](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_sampler_filter_minmax%5D%20%40jeffbolznv%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_sampler_filter_minmax%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2017-05-19

**IP Status**

No known IP claims.

**Contributors**

- Jeff Bolz, NVIDIA
- Piers Daniell, NVIDIA

## [](#_description)Description

In unextended Vulkan, minification and magnification filters such as LINEAR allow sampled image lookups to return a filtered texel value produced by computing a weighted average of a collection of texels in the neighborhood of the texture coordinate provided.

This extension provides a new sampler parameter which allows applications to produce a filtered texel value by computing a component-wise minimum (MIN) or maximum (MAX) of the texels that would normally be averaged. The reduction mode is orthogonal to the minification and magnification filter parameters. The filter parameters are used to identify the set of texels used to produce a final filtered value; the reduction mode identifies how these texels are combined.

## [](#_promotion_to_vulkan_1_2)Promotion to Vulkan 1.2

All functionality in this extension is included in core Vulkan 1.2, with the EXT suffix omitted. The original type, enum, and command names are still available as aliases of the core functionality.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceSamplerFilterMinmaxPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceSamplerFilterMinmaxPropertiesEXT.html)
- Extending [VkSamplerCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerCreateInfo.html):
  
  - [VkSamplerReductionModeCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerReductionModeCreateInfoEXT.html)

## [](#_new_enums)New Enums

- [VkSamplerReductionModeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerReductionModeEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_SAMPLER_FILTER_MINMAX_EXTENSION_NAME`
- `VK_EXT_SAMPLER_FILTER_MINMAX_SPEC_VERSION`
- Extending [VkFormatFeatureFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatFeatureFlagBits.html):
  
  - `VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_MINMAX_BIT_EXT`
- Extending [VkSamplerReductionMode](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerReductionMode.html):
  
  - `VK_SAMPLER_REDUCTION_MODE_MAX_EXT`
  - `VK_SAMPLER_REDUCTION_MODE_MIN_EXT`
  - `VK_SAMPLER_REDUCTION_MODE_WEIGHTED_AVERAGE_EXT`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SAMPLER_FILTER_MINMAX_PROPERTIES_EXT`
  - `VK_STRUCTURE_TYPE_SAMPLER_REDUCTION_MODE_CREATE_INFO_EXT`

## [](#_version_history)Version History

- Revision 2, 2017-05-19 (Piers Daniell)
  
  - Renamed to EXT
- Revision 1, 2017-03-25 (Jeff Bolz)
  
  - Internal revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_sampler_filter_minmax).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0