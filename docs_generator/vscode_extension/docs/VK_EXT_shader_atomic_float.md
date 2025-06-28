# VK\_EXT\_shader\_atomic\_float(3) Manual Page

## Name

VK\_EXT\_shader\_atomic\_float - device extension



## [](#_registered_extension_number)Registered Extension Number

261

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_api_interactions)API Interactions

- Interacts with VkPhysicalDeviceShaderAtomicFloatFeaturesEXT::sparseImageFloat32AtomicAdd
- Interacts with VkPhysicalDeviceShaderAtomicFloatFeaturesEXT::sparseImageFloat32Atomics

## [](#_spir_v_dependencies)SPIR-V Dependencies

- [SPV\_EXT\_shader\_atomic\_float\_add](https://github.khronos.org/SPIRV-Registry/extensions/EXT/SPV_EXT_shader_atomic_float_add.html)

## [](#_contact)Contact

- Vikram Kushwaha [\[GitHub\]vkushwaha-nv](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_shader_atomic_float%5D%20%40vkushwaha-nv%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_shader_atomic_float%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2020-07-15

**IP Status**

No known IP claims.

**Interactions and External Dependencies**

- This extension provides API support for [`GL_EXT_shader_atomic_float`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/ext/GLSL_EXT_shader_atomic_float.txt)

**Contributors**

- Vikram Kushwaha, NVIDIA
- Jeff Bolz, NVIDIA

## [](#_description)Description

This extension allows a shader to contain floating-point atomic operations on buffer, workgroup, and image memory. It also advertises the SPIR-V `AtomicFloat32AddEXT` and `AtomicFloat64AddEXT` capabilities that allows atomic addition on floating-points numbers. The supported operations include `OpAtomicFAddEXT`, `OpAtomicExchange`, `OpAtomicLoad` and `OpAtomicStore`.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceShaderAtomicFloatFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceShaderAtomicFloatFeaturesEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_SHADER_ATOMIC_FLOAT_EXTENSION_NAME`
- `VK_EXT_SHADER_ATOMIC_FLOAT_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_ATOMIC_FLOAT_FEATURES_EXT`

## [](#_new_spir_v_capabilities)New SPIR-V Capabilities

- [`AtomicFloat32AddEXT`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-AtomicFloat32AddEXT)
- [`AtomicFloat64AddEXT`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-AtomicFloat64AddEXT)

## [](#_version_history)Version History

- Revision 1, 2020-07-15 (Vikram Kushwaha)
  
  - Internal revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_shader_atomic_float)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0