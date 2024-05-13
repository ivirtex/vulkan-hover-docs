# VK_NV_shader_sm_builtins(3) Manual Page

## Name

VK_NV_shader_sm_builtins - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

155

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[Version 1.1](#versions-1.1)  

## <a href="#_spir_v_dependencies" class="anchor"></a>SPIR-V Dependencies

- [SPV_NV_shader_sm_builtins](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/NV/SPV_NV_shader_sm_builtins.html)

## <a href="#_contact" class="anchor"></a>Contact

- Daniel Koch <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_shader_sm_builtins%5D%20@dgkoch%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_shader_sm_builtins%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>dgkoch</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2019-05-28

**Interactions and External Dependencies**  
- This extension provides API support for
  [`GL_NV_shader_sm_builtins`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/nv/GLSL_NV_shader_sm_builtins.txt)

**Contributors**  
- Jeff Bolz, NVIDIA

- Eric Werness, NVIDIA

## <a href="#_description" class="anchor"></a>Description

This extension provides the ability to determine device-specific
properties on NVIDIA GPUs. It provides the number of streaming
multiprocessors (SMs), the maximum number of warps (subgroups) that can
run on an SM, and shader builtins to enable invocations to identify
which SM and warp a shader invocation is executing on.

This extension enables support for the SPIR-V `ShaderSMBuiltinsNV`
capability.

These properties and built-ins **should** typically only be used for
debugging purposes.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceShaderSMBuiltinsFeaturesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderSMBuiltinsFeaturesNV.html)

- Extending
  [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html):

  - [VkPhysicalDeviceShaderSMBuiltinsPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderSMBuiltinsPropertiesNV.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_NV_SHADER_SM_BUILTINS_EXTENSION_NAME`

- `VK_NV_SHADER_SM_BUILTINS_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_SM_BUILTINS_FEATURES_NV`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_SM_BUILTINS_PROPERTIES_NV`

## <a href="#_new_or_modified_built_in_variables" class="anchor"></a>New or Modified Built-In Variables

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-warpspersmnv"
  target="_blank" rel="noopener"><code>WarpsPerSMNV</code></a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-smcountnv"
  target="_blank" rel="noopener"><code>SMCountNV</code></a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-warpidnv"
  target="_blank" rel="noopener"><code>WarpIDNV</code></a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-smidnv"
  target="_blank" rel="noopener"><code>SMIDNV</code></a>

## <a href="#_new_spir_v_capabilities" class="anchor"></a>New SPIR-V Capabilities

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#spirvenv-capabilities-table-ShaderSMBuiltinsNV"
  target="_blank" rel="noopener"><code>ShaderSMBuiltinsNV</code></a>

## <a href="#_issues" class="anchor"></a>Issues

1.  What should we call this extension?

    **RESOLVED**: `NV_shader_sm_builtins`. Other options considered
    included:

    - `NV_shader_smid` - but SMID is really easy to typo/confuse as
      SIMD.

    - `NV_shader_sm_info` - but **Info** is typically reserved for input
      structures

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2019-05-28 (Daniel Koch)

  - Internal revisions

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_NV_shader_sm_builtins"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
