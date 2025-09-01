# VK\_KHR\_shader\_untyped\_pointers(3) Manual Page

## Name

VK\_KHR\_shader\_untyped\_pointers - device extension



## [](#_registered_extension_number)Registered Extension Number

388

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)

## [](#_spir_v_dependencies)SPIR-V Dependencies

- [SPV\_KHR\_untyped\_pointers](https://github.khronos.org/SPIRV-Registry/extensions/KHR/SPV_KHR_untyped_pointers.html)

## [](#_contact)Contact

- Alan Baker [\[GitHub\]alan-baker](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_shader_untyped_pointers%5D%20%40alan-baker%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_shader_untyped_pointers%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_KHR\_shader\_untyped\_pointers](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_KHR_shader_untyped_pointers.adoc)

**Last Modified Date**

2024-03-26

**IP Status**

No known IP claims.

**Interactions and External Dependencies**

- Requires the [`SPV_KHR_untyped_pointers`](https://github.khronos.org/SPIRV-Registry/extensions/KHR/SPV_KHR_untyped_pointers.html) SPIR-V extension.

**Contributors**

- Alan Baker, Google
- Jan-Harald Fredriksen, Arm
- Tom Olson, Arm
- Spencer Fricke, LunarG
- Shahbaz Youssefi, Google
- Tobias Hector, AMD

This extension adds Vulkan support for the [`SPV_KHR_untyped_pointers`](https://github.khronos.org/SPIRV-Registry/extensions/KHR/SPV_KHR_untyped_pointers.html) SPIR-V extension. It provides an alternative to strongly-typed pointers. Untyped pointers allow shader authors to reinterpret data accessed through memory and atomic instructions versus the data type declared in the variable without extra conversion instructions. Untyped pointers also provide an efficient translation from templated load/store operations in high-level languages and simplify shaders that support operations, but not storage, on smaller data types (e.g. 16-bit floating-point types).

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceShaderUntypedPointersFeaturesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceShaderUntypedPointersFeaturesKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_SHADER_UNTYPED_POINTERS_EXTENSION_NAME`
- `VK_KHR_SHADER_UNTYPED_POINTERS_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_UNTYPED_POINTERS_FEATURES_KHR`

## [](#_version_history)Version History

- Revision 1, 2024-03-26 (Alan Baker)
  
  - Internal draft version

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_shader_untyped_pointers).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0