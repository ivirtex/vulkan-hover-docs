# VK_QCOM_image_processing(3) Manual Page

## Name

VK_QCOM_image_processing - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

441

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_format_feature_flags2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_format_feature_flags2.html)  
or  
[Version 1.3](#versions-1.3)  

## <a href="#_api_interactions" class="anchor"></a>API Interactions

- Interacts with VK_VERSION_1_3

- Interacts with VK_KHR_format_feature_flags2

## <a href="#_spir_v_dependencies" class="anchor"></a>SPIR-V Dependencies

- [SPV_QCOM_image_processing](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/QCOM/SPV_QCOM_image_processing.html)

## <a href="#_contact" class="anchor"></a>Contact

- Matthew Netsch <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_QCOM_image_processing%5D%20@mnetsch%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_QCOM_image_processing%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>mnetsch</a>

## <a href="#_extension_proposal" class="anchor"></a>Extension Proposal

[VK_QCOM_image_processing](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_QCOM_image_processing.adoc)

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2022-07-08

**Interactions and External Dependencies**  
- This extension provides API support for
  [`GL_QCOM_image_processing`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/qcom/GLSL_QCOM_image_processing.txt)

**Contributors**  
- Jeff Leger, Qualcomm Technologies, Inc.

- Ruihao Zhang, Qualcomm Technologies, Inc.

## <a href="#_description" class="anchor"></a>Description

GPUs are commonly used to process images for various applications from
3D graphics to UI and from composition to compute applications. Simple
scaling and filtering can be done with bilinear filtering, which comes
for free during texture sampling. However, as screen sizes get larger
and more use cases rely on GPU such as camera and video post-processing
needs, there is increasing demand for GPU to support higher order
filtering and other advanced image processing.

This extension introduces a new set of SPIR-V built-in functions for
image processing. It exposes the following new imaging operations

- The `OpImageSampleWeightedQCOM` instruction takes 3 operands: *sampled
  image*, *weight image*, and texture coordinates. The instruction
  computes a weighted average of an MxN region of texels in the *sampled
  image*, using a set of MxN weights in the *weight image*.

- The `OpImageBoxFilterQCOM` instruction takes 3 operands: *sampled
  image*, *box size*, and texture coordinates. Note that *box size*
  specifies a floating-point width and height in texels. The instruction
  computes a weighted average of all texels in the *sampled image* that
  are covered (either partially or fully) by a box with the specified
  size and centered at the specified texture coordinates.

- The `OpImageBlockMatchSADQCOM` and `OpImageBlockMatchSSDQCOM`
  instructions each takes 5 operands: *target image*, *target
  coordinates*, *reference image*, *reference coordinates*, and *block
  size*. Each instruction computes an error metric, that describes
  whether a block of texels in the *target image* matches a
  corresponding block of texels in the *reference image*. The error
  metric is computed per-component. `OpImageBlockMatchSADQCOM` computes
  "Sum Of Absolute Difference" and `OpImageBlockMatchSSDQCOM` computes
  "Sum of Squared Difference".

Each of the image processing instructions operate only on 2D images. The
instructions do not-support sampling of mipmap, multi-plane,
multi-layer, multi-sampled, or depth/stencil images. The instructions
can be used in any shader stage.

Implementations of this extension should support these operations
natively at the HW instruction level, offering potential performance
gains as well as ease of development.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkImageViewCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewCreateInfo.html):

  - [VkImageViewSampleWeightCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewSampleWeightCreateInfoQCOM.html)

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceImageProcessingFeaturesQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageProcessingFeaturesQCOM.html)

- Extending
  [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html):

  - [VkPhysicalDeviceImageProcessingPropertiesQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageProcessingPropertiesQCOM.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_QCOM_IMAGE_PROCESSING_EXTENSION_NAME`

- `VK_QCOM_IMAGE_PROCESSING_SPEC_VERSION`

- Extending [VkDescriptorType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorType.html):

  - `VK_DESCRIPTOR_TYPE_BLOCK_MATCH_IMAGE_QCOM`

  - `VK_DESCRIPTOR_TYPE_SAMPLE_WEIGHT_IMAGE_QCOM`

- Extending [VkImageUsageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageUsageFlagBits.html):

  - `VK_IMAGE_USAGE_SAMPLE_BLOCK_MATCH_BIT_QCOM`

  - `VK_IMAGE_USAGE_SAMPLE_WEIGHT_BIT_QCOM`

- Extending [VkSamplerCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerCreateFlagBits.html):

  - `VK_SAMPLER_CREATE_IMAGE_PROCESSING_BIT_QCOM`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_IMAGE_VIEW_SAMPLE_WEIGHT_CREATE_INFO_QCOM`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_PROCESSING_FEATURES_QCOM`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_PROCESSING_PROPERTIES_QCOM`

If [VK_KHR_format_feature_flags2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_format_feature_flags2.html) or
[Version 1.3](#versions-1.3) is supported:

- Extending [VkFormatFeatureFlagBits2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatFeatureFlagBits2.html):

  - `VK_FORMAT_FEATURE_2_BLOCK_MATCHING_BIT_QCOM`

  - `VK_FORMAT_FEATURE_2_BOX_FILTER_SAMPLED_BIT_QCOM`

  - `VK_FORMAT_FEATURE_2_WEIGHT_IMAGE_BIT_QCOM`

  - `VK_FORMAT_FEATURE_2_WEIGHT_SAMPLED_IMAGE_BIT_QCOM`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2022-07-08 (Jeff Leger)

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_QCOM_image_processing"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
