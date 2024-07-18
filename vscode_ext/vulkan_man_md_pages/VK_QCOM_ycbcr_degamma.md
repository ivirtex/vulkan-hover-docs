# VK_QCOM_ycbcr_degamma(3) Manual Page

## Name

VK_QCOM_ycbcr_degamma - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

521

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

None

## <a href="#_contact" class="anchor"></a>Contact

- Matthew Netsch <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_QCOM_ycbcr_degamma%5D%20@mnetsch%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_QCOM_ycbcr_degamma%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>mnetsch</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2023-07-31

**IP Status**  
No known IP claims.

**Interactions and External Dependencies**  
None

**Contributors**  
- Jeff Leger, Qualcomm

- Jonathan Wicks, Qualcomm

## <a href="#_description" class="anchor"></a>Description

This extension allows implementations to expose support for “sRGB EOTF”
also known as “sRGB degamma”, used in combination with images using
8-bit Y′C<sub>B</sub>C<sub>R</sub> formats. In addition, the degamma can
be selectively applied to the Y (luma) or CrCb (chroma).

[`VK_KHR_sampler_ycbcr_conversion`](VK_KHR_sampler_ycbcr_conversion.html)
adds support for Y′C<sub>B</sub>C<sub>R</sub> conversion, but allows
texture sampling in a non-linear space which can cause artifacts. This
extension allows implementations to expose sRGB degamma for
Y′C<sub>B</sub>C<sub>R</sub> formats, which is performed during texture
filtering, allowing texture filtering to operate in a linear space.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceYcbcrDegammaFeaturesQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceYcbcrDegammaFeaturesQCOM.html)

- Extending
  [VkSamplerYcbcrConversionCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrConversionCreateInfo.html):

  - [VkSamplerYcbcrConversionYcbcrDegammaCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrConversionYcbcrDegammaCreateInfoQCOM.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_QCOM_YCBCR_DEGAMMA_EXTENSION_NAME`

- `VK_QCOM_YCBCR_DEGAMMA_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_YCBCR_DEGAMMA_FEATURES_QCOM`

  - `VK_STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_YCBCR_DEGAMMA_CREATE_INFO_QCOM`

## <a href="#_issues" class="anchor"></a>Issues

1\) Which Y′C<sub>B</sub>C<sub>R</sub> formats support the degamma
feature?

**RESOLVED**: For implementations that support the extension, each
format that contains 8-bit R, G, and B components and supports either
`VK_FORMAT_FEATURE_MIDPOINT_CHROMA_SAMPLES_BIT` or
`VK_FORMAT_FEATURE_COSITED_CHROMA_SAMPLES_BIT` must support degamma.

Since non-compressed Vulkan sRGB formats are already limited to 8-bit
components, and since Adreno supports degamma for all 8bit
Y′C<sub>B</sub>C<sub>R</sub> formats, this extension does not introduce
a new VK_FORMAT_FEATURE\* bit for the degamma feature.

2\) On which Y′C<sub>B</sub>C<sub>R</sub> components is the degamma
applied?

**RESOLVED**: While degamma is expected to be applied to only the Y
(luma) component, the extension provides the ability to selectively
enable degamma for both the Y (luma) and/or CbCr (chroma) components.

3\) Should degamma be enabled for the sampler object or for the image
view object?

**RESOLVED**: Both. This extension extends
[VkSamplerYcbcrConversionCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrConversionCreateInfo.html)
and the specification already requires that both sampler and view
objects must be created with an *identical*
[VkSamplerYcbcrConversionCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrConversionCreateInfo.html)
in their pNext chains.

4\) Why apply the “sRGB” transfer function directly to
Y′C<sub>B</sub>C<sub>R</sub> data when it would be more correct to use
the “ITU transfer function”, and do so only after the values have been
converted into non-linear R’G’B'?

**RESOLVED**: Y′C<sub>B</sub>C<sub>R</sub> is frequently stored
according to standards (e.g. BT.601 and BT.709) that specify that the
conversion between linear and non-linear should use the ITU Transfer
function. The ITU transfer function is mathematically different from the
sRGB transfer function and while sRGB and ITU define similar curves, the
difference is significant. Performing the “sRGB degamma” prior to range
expansion can introduce artifacts if the content uses
`VK_SAMPLER_YCBCR_RANGE_ITU_NARROW` encoding. Nevertheless, using sRGB
can make sense for certain use cases where camera YCbCr images are known
to be encoded with sRGB (or a pure gamma 2.2) transfer function and are
known to use full-range encoding.

For those use cases, this extension leverages the GPU ability to enable
sRGB degamma at little cost, and can improve quality because texture
filtering is able to occur in linear space.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2023-07-31 (Jeff Leger)

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_QCOM_ycbcr_degamma"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
