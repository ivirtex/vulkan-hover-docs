# VK\_QCOM\_ycbcr\_degamma(3) Manual Page

## Name

VK\_QCOM\_ycbcr\_degamma - device extension



## [](#_registered_extension_number)Registered Extension Number

521

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

None

## [](#_contact)Contact

- Matthew Netsch [\[GitHub\]mnetsch](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_QCOM_ycbcr_degamma%5D%20%40mnetsch%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_QCOM_ycbcr_degamma%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2023-07-31

**IP Status**

No known IP claims.

**Interactions and External Dependencies**

None

**Contributors**

- Jeff Leger, Qualcomm
- Jonathan Wicks, Qualcomm

## [](#_description)Description

This extension allows implementations to expose support for “sRGB EOTF” also known as “sRGB degamma”, used in combination with images using 8-bit Y′CBCR formats. In addition, the degamma can be selectively applied to the Y (luma) or CrCb (chroma).

`VK_KHR_sampler_ycbcr_conversion` adds support for Y′CBCR conversion, but allows texture sampling in a non-linear space which can cause artifacts. This extension allows implementations to expose sRGB degamma for Y′CBCR formats, which is performed during texture filtering, allowing texture filtering to operate in a linear space.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceYcbcrDegammaFeaturesQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceYcbcrDegammaFeaturesQCOM.html)
- Extending [VkSamplerYcbcrConversionCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrConversionCreateInfo.html):
  
  - [VkSamplerYcbcrConversionYcbcrDegammaCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrConversionYcbcrDegammaCreateInfoQCOM.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_QCOM_YCBCR_DEGAMMA_EXTENSION_NAME`
- `VK_QCOM_YCBCR_DEGAMMA_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_YCBCR_DEGAMMA_FEATURES_QCOM`
  - `VK_STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_YCBCR_DEGAMMA_CREATE_INFO_QCOM`

## [](#_issues)Issues

1\) Which Y′CBCR formats support the degamma feature?

**RESOLVED**: For implementations that support the extension, each format that contains 8-bit R, G, and B components and supports either `VK_FORMAT_FEATURE_MIDPOINT_CHROMA_SAMPLES_BIT` or `VK_FORMAT_FEATURE_COSITED_CHROMA_SAMPLES_BIT` must support degamma.

Since non-compressed Vulkan sRGB formats are already limited to 8-bit components, and since Adreno supports degamma for all 8bit Y′CBCR formats, this extension does not introduce a new VK\_FORMAT\_FEATURE* bit for the degamma feature.

2\) On which Y′CBCR components is the degamma applied?

**RESOLVED**: While degamma is expected to be applied to only the Y (luma) component, the extension provides the ability to selectively enable degamma for both the Y (luma) and/or CbCr (chroma) components.

3\) Should degamma be enabled for the sampler object or for the image view object?

**RESOLVED**: Both. This extension extends [VkSamplerYcbcrConversionCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrConversionCreateInfo.html) and the specification already requires that both sampler and view objects must be created with an *identical* [VkSamplerYcbcrConversionCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrConversionCreateInfo.html) in their pNext chains.

4\) Why apply the “sRGB” transfer function directly to Y′CBCR data when it would be more correct to use the “ITU transfer function”, and do so only after the values have been converted into non-linear R’G’B'?

**RESOLVED**: Y′CBCR is frequently stored according to standards (e.g. BT.601 and BT.709) that specify that the conversion between linear and non-linear should use the ITU Transfer function. The ITU transfer function is mathematically different from the sRGB transfer function and while sRGB and ITU define similar curves, the difference is significant. Performing the “sRGB degamma” prior to range expansion can introduce artifacts if the content uses `VK_SAMPLER_YCBCR_RANGE_ITU_NARROW` encoding. Nevertheless, using sRGB can make sense for certain use cases where camera YCbCr images are known to be encoded with sRGB (or a pure gamma 2.2) transfer function and are known to use full-range encoding.

For those use cases, this extension leverages the GPU ability to enable sRGB degamma at little cost, and can improve quality because texture filtering is able to occur in linear space.

## [](#_version_history)Version History

- Revision 1, 2023-07-31 (Jeff Leger)

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_QCOM_ycbcr_degamma)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0