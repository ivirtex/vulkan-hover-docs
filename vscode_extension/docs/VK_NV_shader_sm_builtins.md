# VK\_NV\_shader\_sm\_builtins(3) Manual Page

## Name

VK\_NV\_shader\_sm\_builtins - device extension



## [](#_registered_extension_number)Registered Extension Number

155

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[Vulkan Version 1.1](#versions-1.1)

## [](#_spir_v_dependencies)SPIR-V Dependencies

- [SPV\_NV\_shader\_sm\_builtins](https://github.khronos.org/SPIRV-Registry/extensions/NV/SPV_NV_shader_sm_builtins.html)

## [](#_contact)Contact

- Daniel Koch [\[GitHub\]dgkoch](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_shader_sm_builtins%5D%20%40dgkoch%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_shader_sm_builtins%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2019-05-28

**Interactions and External Dependencies**

- This extension provides API support for [`GL_NV_shader_sm_builtins`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/nv/GLSL_NV_shader_sm_builtins.txt)

**Contributors**

- Jeff Bolz, NVIDIA
- Eric Werness, NVIDIA

## [](#_description)Description

This extension provides the ability to determine device-specific properties on NVIDIA GPUs. It provides the number of streaming multiprocessors (SMs), the maximum number of warps (subgroups) that can run on an SM, and shader builtins to enable invocations to identify which SM and warp a shader invocation is executing on.

This extension enables support for the SPIR-V `ShaderSMBuiltinsNV` capability.

These properties and built-ins **should** typically only be used for debugging purposes.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceShaderSMBuiltinsFeaturesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceShaderSMBuiltinsFeaturesNV.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceShaderSMBuiltinsPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceShaderSMBuiltinsPropertiesNV.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_NV_SHADER_SM_BUILTINS_EXTENSION_NAME`
- `VK_NV_SHADER_SM_BUILTINS_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_SM_BUILTINS_FEATURES_NV`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_SM_BUILTINS_PROPERTIES_NV`

## [](#_new_or_modified_built_in_variables)New or Modified Built-In Variables

- [`WarpsPerSMNV`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-warpspersmnv)
- [`SMCountNV`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-smcountnv)
- [`WarpIDNV`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-warpidnv)
- [`SMIDNV`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-smidnv)

## [](#_new_spir_v_capabilities)New SPIR-V Capabilities

- [`ShaderSMBuiltinsNV`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-ShaderSMBuiltinsNV)

## [](#_issues)Issues

1. What should we call this extension?
   
   **RESOLVED**: `NV_shader_sm_builtins`. Other options considered included:
   
   - `NV_shader_smid` - but SMID is really easy to typo/confuse as SIMD.
   - `NV_shader_sm_info` - but **Info** is typically reserved for input structures

## [](#_version_history)Version History

- Revision 1, 2019-05-28 (Daniel Koch)
  
  - Internal revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_NV_shader_sm_builtins)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0