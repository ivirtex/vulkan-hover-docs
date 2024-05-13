# VK_NV_cooperative_matrix(3) Manual Page

## Name

VK_NV_cooperative_matrix - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

250

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_spir_v_dependencies" class="anchor"></a>SPIR-V Dependencies

- [SPV_NV_cooperative_matrix](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/NV/SPV_NV_cooperative_matrix.html)

## <a href="#_contact" class="anchor"></a>Contact

- Jeff Bolz <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_cooperative_matrix%5D%20@jeffbolznv%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_cooperative_matrix%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>jeffbolznv</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2019-02-05

**Interactions and External Dependencies**  
- This extension provides API support for
  [`GL_NV_cooperative_matrix`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/nv/GLSL_NV_cooperative_matrix.txt)

**Contributors**  
- Jeff Bolz, NVIDIA

- Markus Tavenrath, NVIDIA

- Daniel Koch, NVIDIA

## <a href="#_description" class="anchor"></a>Description

This extension adds support for using cooperative matrix types in
SPIR-V. Cooperative matrix types are medium-sized matrices that are
primarily supported in compute shaders, where the storage for the matrix
is spread across all invocations in some scope (usually a subgroup) and
those invocations cooperate to efficiently perform matrix multiplies.

Cooperative matrix types are defined by the
[`SPV_NV_cooperative_matrix`](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/NV/SPV_NV_cooperative_matrix.html)
SPIR-V extension and can be used with the
[`GL_NV_cooperative_matrix`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/nv/GLSL_NV_cooperative_matrix.txt)
GLSL extension.

This extension includes support for enumerating the matrix types and
dimensions that are supported by the implementation.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkGetPhysicalDeviceCooperativeMatrixPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceCooperativeMatrixPropertiesNV.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkCooperativeMatrixPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCooperativeMatrixPropertiesNV.html)

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceCooperativeMatrixFeaturesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceCooperativeMatrixFeaturesNV.html)

- Extending
  [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html):

  - [VkPhysicalDeviceCooperativeMatrixPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceCooperativeMatrixPropertiesNV.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkComponentTypeNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkComponentTypeNV.html)

- [VkScopeNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkScopeNV.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_NV_COOPERATIVE_MATRIX_EXTENSION_NAME`

- `VK_NV_COOPERATIVE_MATRIX_SPEC_VERSION`

- Extending [VkComponentTypeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkComponentTypeKHR.html):

  - `VK_COMPONENT_TYPE_FLOAT16_NV`

  - `VK_COMPONENT_TYPE_FLOAT32_NV`

  - `VK_COMPONENT_TYPE_FLOAT64_NV`

  - `VK_COMPONENT_TYPE_SINT16_NV`

  - `VK_COMPONENT_TYPE_SINT32_NV`

  - `VK_COMPONENT_TYPE_SINT64_NV`

  - `VK_COMPONENT_TYPE_SINT8_NV`

  - `VK_COMPONENT_TYPE_UINT16_NV`

  - `VK_COMPONENT_TYPE_UINT32_NV`

  - `VK_COMPONENT_TYPE_UINT64_NV`

  - `VK_COMPONENT_TYPE_UINT8_NV`

- Extending [VkScopeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkScopeKHR.html):

  - `VK_SCOPE_DEVICE_NV`

  - `VK_SCOPE_QUEUE_FAMILY_NV`

  - `VK_SCOPE_SUBGROUP_NV`

  - `VK_SCOPE_WORKGROUP_NV`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_COOPERATIVE_MATRIX_PROPERTIES_NV`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_COOPERATIVE_MATRIX_FEATURES_NV`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_COOPERATIVE_MATRIX_PROPERTIES_NV`

## <a href="#_new_spir_v_capabilities" class="anchor"></a>New SPIR-V Capabilities

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#spirvenv-capabilities-table-CooperativeMatrixNV"
  target="_blank" rel="noopener"><code>CooperativeMatrixNV</code></a>

## <a href="#_issues" class="anchor"></a>Issues

\(1\) What matrix properties will be supported in practice?

**RESOLVED**: In NVIDIA’s initial implementation, we will support:

- AType = BType = fp16 CType = DType = fp16 MxNxK = 16x8x16 scope =
  Subgroup

- AType = BType = fp16 CType = DType = fp16 MxNxK = 16x8x8 scope =
  Subgroup

- AType = BType = fp16 CType = DType = fp32 MxNxK = 16x8x16 scope =
  Subgroup

- AType = BType = fp16 CType = DType = fp32 MxNxK = 16x8x8 scope =
  Subgroup

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2019-02-05 (Jeff Bolz)

  - Internal revisions

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_NV_cooperative_matrix"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
