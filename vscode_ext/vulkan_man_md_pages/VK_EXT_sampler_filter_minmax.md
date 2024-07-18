# VK_EXT_sampler_filter_minmax(3) Manual Page

## Name

VK_EXT_sampler_filter_minmax - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

131

## <a href="#_revision" class="anchor"></a>Revision

2

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_deprecation_state" class="anchor"></a>Deprecation State

- *Promoted* to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#versions-1.2-promotions"
  target="_blank" rel="noopener">Vulkan 1.2</a>

## <a href="#_contact" class="anchor"></a>Contact

- Jeff Bolz <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_sampler_filter_minmax%5D%20@jeffbolznv%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_sampler_filter_minmax%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>jeffbolznv</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2017-05-19

**IP Status**  
No known IP claims.

**Contributors**  
- Jeff Bolz, NVIDIA

- Piers Daniell, NVIDIA

## <a href="#_description" class="anchor"></a>Description

In unextended Vulkan, minification and magnification filters such as
LINEAR allow sampled image lookups to return a filtered texel value
produced by computing a weighted average of a collection of texels in
the neighborhood of the texture coordinate provided.

This extension provides a new sampler parameter which allows
applications to produce a filtered texel value by computing a
component-wise minimum (MIN) or maximum (MAX) of the texels that would
normally be averaged. The reduction mode is orthogonal to the
minification and magnification filter parameters. The filter parameters
are used to identify the set of texels used to produce a final filtered
value; the reduction mode identifies how these texels are combined.

## <a href="#_promotion_to_vulkan_1_2" class="anchor"></a>Promotion to Vulkan 1.2

All functionality in this extension is included in core Vulkan 1.2, with
the EXT suffix omitted. The original type, enum and command names are
still available as aliases of the core functionality.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending
  [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html):

  - [VkPhysicalDeviceSamplerFilterMinmaxPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSamplerFilterMinmaxPropertiesEXT.html)

- Extending [VkSamplerCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerCreateInfo.html):

  - [VkSamplerReductionModeCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerReductionModeCreateInfoEXT.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkSamplerReductionModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerReductionModeEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_SAMPLER_FILTER_MINMAX_EXTENSION_NAME`

- `VK_EXT_SAMPLER_FILTER_MINMAX_SPEC_VERSION`

- Extending [VkFormatFeatureFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatFeatureFlagBits.html):

  - `VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_MINMAX_BIT_EXT`

- Extending [VkSamplerReductionMode](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerReductionMode.html):

  - `VK_SAMPLER_REDUCTION_MODE_MAX_EXT`

  - `VK_SAMPLER_REDUCTION_MODE_MIN_EXT`

  - `VK_SAMPLER_REDUCTION_MODE_WEIGHTED_AVERAGE_EXT`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SAMPLER_FILTER_MINMAX_PROPERTIES_EXT`

  - `VK_STRUCTURE_TYPE_SAMPLER_REDUCTION_MODE_CREATE_INFO_EXT`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 2, 2017-05-19 (Piers Daniell)

  - Renamed to EXT

- Revision 1, 2017-03-25 (Jeff Bolz)

  - Internal revisions

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_sampler_filter_minmax"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
