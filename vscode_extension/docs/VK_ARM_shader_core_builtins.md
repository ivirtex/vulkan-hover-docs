# VK\_ARM\_shader\_core\_builtins(3) Manual Page

## Name

VK\_ARM\_shader\_core\_builtins - device extension



## [](#_registered_extension_number)Registered Extension Number

498

## [](#_revision)Revision

2

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_spir_v_dependencies)SPIR-V Dependencies

- [SPV\_ARM\_core\_builtins](https://github.khronos.org/SPIRV-Registry/extensions/ARM/SPV_ARM_core_builtins.html)

## [](#_contact)Contact

- Kevin Petit [\[GitHub\]kpet](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_ARM_shader_core_builtins%5D%20%40kpet%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_ARM_shader_core_builtins%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2022-10-05

**Interactions and External Dependencies**

- This extension provides API support for [`GL_ARM_shader_core_builtins`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/arm/GLSL_ARM_shader_core_builtins.txt)

**Contributors**

- Kevin Petit, Arm Ltd.
- Jan-Harald Fredriksen, Arm Ltd.

## [](#_description)Description

This extension provides the ability to determine device-specific properties on Arm GPUs. It exposes properties for the number of shader cores, the maximum number of warps that can run on a shader core, and shader builtins to enable invocations to identify which core and warp a shader invocation is executing on.

This extension enables support for the SPIR-V `CoreBuiltinsARM` capability.

These properties and built-ins can be used for debugging or performance optimization purposes. A typical optimization example would be to use `CoreIDARM` to select a per-shader-core instance of a data structure in algorithms that use atomics so as to reduce contention.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceShaderCoreBuiltinsFeaturesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceShaderCoreBuiltinsFeaturesARM.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceShaderCoreBuiltinsPropertiesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceShaderCoreBuiltinsPropertiesARM.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_ARM_SHADER_CORE_BUILTINS_EXTENSION_NAME`
- `VK_ARM_SHADER_CORE_BUILTINS_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_CORE_BUILTINS_FEATURES_ARM`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_CORE_BUILTINS_PROPERTIES_ARM`

## [](#_new_or_modified_built_in_variables)New or Modified Built-In Variables

- [`CoreCountARM`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-corecountarm)
- [`CoreMaxIDARM`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-coremaxidarm)
- [`CoreIDARM`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-coreidarm)
- [`WarpsMaxIDARM`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-warpmaxidarm)
- [`WarpIDARM`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-warpidarm)

## [](#_new_spir_v_capabilities)New SPIR-V Capabilities

- [`CoreBuiltinsARM`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-CoreBuiltinsARM)

## [](#_issues)Issues

None.

## [](#_version_history)Version History

- Revision 1, 2022-10-05 (Kevin Petit)
  
  - Initial revision
- Revision 2, 2022-10-26 (Kevin Petit)
  
  - Add `shaderCoreMask` property

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_ARM_shader_core_builtins)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0