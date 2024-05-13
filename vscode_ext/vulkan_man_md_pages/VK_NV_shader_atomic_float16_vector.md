# VK_NV_shader_atomic_float16_vector(3) Manual Page

## Name

VK_NV_shader_atomic_float16_vector - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

564

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

None

## <a href="#_spir_v_dependencies" class="anchor"></a>SPIR-V Dependencies

- [SPV_NV_shader_atomic_fp16_vector](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/NV/SPV_NV_shader_atomic_fp16_vector.html)

## <a href="#_contact" class="anchor"></a>Contact

- Jeff Bolz <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_shader_atomic_float16_vector%5D%20@jeffbolznv%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_shader_atomic_float16_vector%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>jeffbolznv</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2024-02-03

**IP Status**  
No known IP claims.

**Interactions and External Dependencies**  
- This extension provides API support for
  [`GL_NV_shader_atomic_fp16_vector`](https://registry.khronos.org/OpenGL/extensions/NV/NV_shader_atomic_fp16_vector.txt)

**Contributors**  
- Jeff Bolz, NVIDIA

## <a href="#_description" class="anchor"></a>Description

This extension allows a shader to perform atomic add, min, max, and
exchange operations on 2- and 4-component vectors of float16. Buffer,
workgroup, and image storage classes are all supported.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceShaderAtomicFloat16VectorFeaturesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderAtomicFloat16VectorFeaturesNV.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_NV_SHADER_ATOMIC_FLOAT16_VECTOR_EXTENSION_NAME`

- `VK_NV_SHADER_ATOMIC_FLOAT16_VECTOR_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_ATOMIC_FLOAT16_VECTOR_FEATURES_NV`

## <a href="#_issues" class="anchor"></a>Issues

None.

## <a href="#_new_spir_v_capabilities" class="anchor"></a>New SPIR-V Capabilities

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#spirvenv-capabilities-table-AtomicFloat16VectorNV"
  target="_blank" rel="noopener"><code>AtomicFloat16VectorNV</code></a>

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2024-02-03 (Jeff Bolz)

  - Internal revisions

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_NV_shader_atomic_float16_vector"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
