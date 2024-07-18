# VK_ARM_shader_core_builtins(3) Manual Page

## Name

VK_ARM_shader_core_builtins - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

498

## <a href="#_revision" class="anchor"></a>Revision

2

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_spir_v_dependencies" class="anchor"></a>SPIR-V Dependencies

- [SPV_ARM_core_builtins](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/ARM/SPV_ARM_core_builtins.html)

## <a href="#_contact" class="anchor"></a>Contact

- Kevin Petit <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_ARM_shader_core_builtins%5D%20@kpet%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_ARM_shader_core_builtins%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>kpet</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2022-10-05

**Interactions and External Dependencies**  
- This extension provides API support for
  [`GL_ARM_shader_core_builtins`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/arm/GLSL_ARM_shader_core_builtins.txt)

**Contributors**  
- Kevin Petit, Arm Ltd.

- Jan-Harald Fredriksen, Arm Ltd.

## <a href="#_description" class="anchor"></a>Description

This extension provides the ability to determine device-specific
properties on Arm GPUs. It exposes properties for the number of shader
cores, the maximum number of warps that can run on a shader core, and
shader builtins to enable invocations to identify which core and warp a
shader invocation is executing on.

This extension enables support for the SPIR-V `CoreBuiltinsARM`
capability.

These properties and built-ins can be used for debugging or performance
optimization purposes. A typical optimization example would be to use
`CoreIDARM` to select a per-shader-core instance of a data structure in
algorithms that use atomics so as to reduce contention.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceShaderCoreBuiltinsFeaturesARM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderCoreBuiltinsFeaturesARM.html)

- Extending
  [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html):

  - [VkPhysicalDeviceShaderCoreBuiltinsPropertiesARM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderCoreBuiltinsPropertiesARM.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_ARM_SHADER_CORE_BUILTINS_EXTENSION_NAME`

- `VK_ARM_SHADER_CORE_BUILTINS_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_CORE_BUILTINS_FEATURES_ARM`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_CORE_BUILTINS_PROPERTIES_ARM`

## <a href="#_new_or_modified_built_in_variables" class="anchor"></a>New or Modified Built-In Variables

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-corecountarm"
  target="_blank" rel="noopener"><code>CoreCountARM</code></a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-coremaxidarm"
  target="_blank" rel="noopener"><code>CoreMaxIDARM</code></a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-coreidarm"
  target="_blank" rel="noopener"><code>CoreIDARM</code></a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-warpmaxidarm"
  target="_blank" rel="noopener"><code>WarpsMaxIDARM</code></a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-warpidarm"
  target="_blank" rel="noopener"><code>WarpIDARM</code></a>

## <a href="#_new_spir_v_capabilities" class="anchor"></a>New SPIR-V Capabilities

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#spirvenv-capabilities-table-CoreBuiltinsARM"
  target="_blank" rel="noopener"><code>CoreBuiltinsARM</code></a>

## <a href="#_issues" class="anchor"></a>Issues

None.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2022-10-05 (Kevin Petit)

  - Initial revision

- Revision 2, 2022-10-26 (Kevin Petit)

  - Add `shaderCoreMask` property

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_ARM_shader_core_builtins"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
