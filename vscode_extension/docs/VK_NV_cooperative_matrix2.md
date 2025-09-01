# VK\_NV\_cooperative\_matrix2(3) Manual Page

## Name

VK\_NV\_cooperative\_matrix2 - device extension



## [](#_registered_extension_number)Registered Extension Number

594

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_cooperative\_matrix](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_cooperative_matrix.html)

## [](#_spir_v_dependencies)SPIR-V Dependencies

- [SPV\_NV\_cooperative\_matrix2](https://github.khronos.org/SPIRV-Registry/extensions/NV/SPV_NV_cooperative_matrix2.html)
- [SPV\_NV\_tensor\_addressing](https://github.khronos.org/SPIRV-Registry/extensions/NV/SPV_NV_tensor_addressing.html)

## [](#_contact)Contact

- Jeff Bolz [\[GitHub\]jeffbolznv](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_cooperative_matrix2%5D%20%40jeffbolznv%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_cooperative_matrix2%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_NV\_cooperative\_matrix2](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_NV_cooperative_matrix2.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2024-08-01

**Interactions and External Dependencies**

- This extension provides API support for [`GLSL_NV_cooperative_matrix2`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/nv/GLSL_NV_cooperative_matrix2.txt)

**Contributors**

- Jeff Bolz, NVIDIA
- Karthik Vaidyanathan, NVIDIA

## [](#_description)Description

This extension adds several new features building on the cooperative matrix types added in VK\_KHR\_cooperative\_matrix. The goal is to add and accelerate features beyond just simple GEMM kernels, including adding support for type/use conversions, reductions, per-element operations, and tensor addressing, and also to improve usability and out-of-the-box performance by adding support for more flexible matrix sizes, and workgroup scope matrices with compiler-managed staging through shared memory.

The new functionality is defined by the [`SPV_NV_tensor_addressing`](https://github.khronos.org/SPIRV-Registry/extensions/NV/SPV_NV_tensor_addressing.html) and [`SPV_NV_cooperative_matrix2`](https://github.khronos.org/SPIRV-Registry/extensions/NV/SPV_NV_cooperative_matrix2.html) SPIR-V extensions and can be used with the [`GLSL_NV_cooperative_matrix2`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/nv/GLSL_NV_cooperative_matrix2.txt) GLSL extension.

This extension includes support for enumerating the matrix types and dimensions that are supported by the implementation, and which specific features are supported.

## [](#_new_commands)New Commands

- [vkGetPhysicalDeviceCooperativeMatrixFlexibleDimensionsPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceCooperativeMatrixFlexibleDimensionsPropertiesNV.html)

## [](#_new_structures)New Structures

- [VkCooperativeMatrixFlexibleDimensionsPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCooperativeMatrixFlexibleDimensionsPropertiesNV.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceCooperativeMatrix2FeaturesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceCooperativeMatrix2FeaturesNV.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceCooperativeMatrix2PropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceCooperativeMatrix2PropertiesNV.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_NV_COOPERATIVE_MATRIX_2_EXTENSION_NAME`
- `VK_NV_COOPERATIVE_MATRIX_2_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_COOPERATIVE_MATRIX_FLEXIBLE_DIMENSIONS_PROPERTIES_NV`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_COOPERATIVE_MATRIX_2_FEATURES_NV`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_COOPERATIVE_MATRIX_2_PROPERTIES_NV`

## [](#_new_spir_v_capabilities)New SPIR-V Capabilities

- [TensorAddressingNV](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-TensorAddressingNV)
- [CooperativeMatrixReductionsNV](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-CooperativeMatrixReductionsNV)
- [CooperativeMatrixConversionsNV](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-CooperativeMatrixConversionsNV)
- [CooperativeMatrixPerElementOperationsNV](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-CooperativeMatrixPerElementOperationsNV)
- [CooperativeMatrixTensorAddressingNV](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-CooperativeMatrixTensorAddressingNV)
- [CooperativeMatrixBlockLoadsNV](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-CooperativeMatrixBlockLoadsNV)

## [](#_issues)Issues

## [](#_version_history)Version History

- Revision 1, 2024-08-01 (Jeff Bolz)
  
  - Initial revision

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_NV_cooperative_matrix2).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0