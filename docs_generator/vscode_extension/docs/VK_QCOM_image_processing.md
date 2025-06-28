# VK\_QCOM\_image\_processing(3) Manual Page

## Name

VK\_QCOM\_image\_processing - device extension



## [](#_registered_extension_number)Registered Extension Number

441

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_format\_feature\_flags2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_format_feature_flags2.html)  
or  
[Vulkan Version 1.3](#versions-1.3)

## [](#_api_interactions)API Interactions

- Interacts with VK\_VERSION\_1\_3
- Interacts with VK\_KHR\_format\_feature\_flags2

## [](#_spir_v_dependencies)SPIR-V Dependencies

- [SPV\_QCOM\_image\_processing](https://github.khronos.org/SPIRV-Registry/extensions/QCOM/SPV_QCOM_image_processing.html)

## [](#_contact)Contact

- Matthew Netsch [\[GitHub\]mnetsch](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_QCOM_image_processing%5D%20%40mnetsch%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_QCOM_image_processing%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_QCOM\_image\_processing](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_QCOM_image_processing.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2022-07-08

**Interactions and External Dependencies**

- This extension provides API support for [`GL_QCOM_image_processing`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/qcom/GLSL_QCOM_image_processing.txt)

**Contributors**

- Jeff Leger, Qualcomm Technologies, Inc.
- Ruihao Zhang, Qualcomm Technologies, Inc.

## [](#_description)Description

GPUs are commonly used to process images for various applications from 3D graphics to UI and from composition to compute applications. Simple scaling and filtering can be done with bilinear filtering, which comes for free during texture sampling. However, as screen sizes get larger and more use cases rely on GPU such as camera and video post-processing needs, there is increasing demand for GPU to support higher order filtering and other advanced image processing.

This extension introduces a new set of SPIR-V built-in functions for image processing. It exposes the following new imaging operations

- The `OpImageSampleWeightedQCOM` instruction takes 3 operands: *sampled image*, *weight image*, and texture coordinates. The instruction computes a weighted average of an MxN region of texels in the *sampled image*, using a set of MxN weights in the *weight image*.
- The `OpImageBoxFilterQCOM` instruction takes 3 operands: *sampled image*, *box size*, and texture coordinates. Note that *box size* specifies a floating-point width and height in texels. The instruction computes a weighted average of all texels in the *sampled image* that are covered (either partially or fully) by a box with the specified size and centered at the specified texture coordinates.
- The `OpImageBlockMatchSADQCOM` and `OpImageBlockMatchSSDQCOM` instructions each takes 5 operands: *target image*, *target coordinates*, *reference image*, *reference coordinates*, and *block size*. Each instruction computes an error metric, that describes whether a block of texels in the *target image* matches a corresponding block of texels in the *reference image*. The error metric is computed per-component. `OpImageBlockMatchSADQCOM` computes "Sum Of Absolute Difference" and `OpImageBlockMatchSSDQCOM` computes "Sum of Squared Difference".

Each of the image processing instructions operate only on 2D images. The instructions do not-support sampling of mipmap, multi-plane, multi-layer, multi-sampled, or depth/stencil images. The instructions can be used in any shader stage.

Implementations of this extension should support these operations natively at the HW instruction level, offering potential performance gains as well as ease of development.

## [](#_new_structures)New Structures

- Extending [VkImageViewCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewCreateInfo.html):
  
  - [VkImageViewSampleWeightCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewSampleWeightCreateInfoQCOM.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceImageProcessingFeaturesQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceImageProcessingFeaturesQCOM.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceImageProcessingPropertiesQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceImageProcessingPropertiesQCOM.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_QCOM_IMAGE_PROCESSING_EXTENSION_NAME`
- `VK_QCOM_IMAGE_PROCESSING_SPEC_VERSION`
- Extending [VkDescriptorType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorType.html):
  
  - `VK_DESCRIPTOR_TYPE_BLOCK_MATCH_IMAGE_QCOM`
  - `VK_DESCRIPTOR_TYPE_SAMPLE_WEIGHT_IMAGE_QCOM`
- Extending [VkImageUsageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageUsageFlagBits.html):
  
  - `VK_IMAGE_USAGE_SAMPLE_BLOCK_MATCH_BIT_QCOM`
  - `VK_IMAGE_USAGE_SAMPLE_WEIGHT_BIT_QCOM`
- Extending [VkSamplerCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerCreateFlagBits.html):
  
  - `VK_SAMPLER_CREATE_IMAGE_PROCESSING_BIT_QCOM`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_IMAGE_VIEW_SAMPLE_WEIGHT_CREATE_INFO_QCOM`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_PROCESSING_FEATURES_QCOM`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_PROCESSING_PROPERTIES_QCOM`

If [VK\_KHR\_format\_feature\_flags2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_format_feature_flags2.html) or [Vulkan Version 1.3](#versions-1.3) is supported:

- Extending [VkFormatFeatureFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatFeatureFlagBits2.html):
  
  - `VK_FORMAT_FEATURE_2_BLOCK_MATCHING_BIT_QCOM`
  - `VK_FORMAT_FEATURE_2_BOX_FILTER_SAMPLED_BIT_QCOM`
  - `VK_FORMAT_FEATURE_2_WEIGHT_IMAGE_BIT_QCOM`
  - `VK_FORMAT_FEATURE_2_WEIGHT_SAMPLED_IMAGE_BIT_QCOM`

## [](#_version_history)Version History

- Revision 1, 2022-07-08 (Jeff Leger)

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_QCOM_image_processing)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0