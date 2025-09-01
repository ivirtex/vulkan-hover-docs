# vkUpdateIndirectExecutionSetShaderEXT(3) Manual Page

## Name

vkUpdateIndirectExecutionSetShaderEXT - Update the contents of an indirect execution set



## [](#_c_specification)C Specification

Shader object elements in an Indirect Execution Set can be updated by calling:

```c++
// Provided by VK_EXT_device_generated_commands
void vkUpdateIndirectExecutionSetShaderEXT(
    VkDevice                                    device,
    VkIndirectExecutionSetEXT                   indirectExecutionSet,
    uint32_t                                    executionSetWriteCount,
    const VkWriteIndirectExecutionSetShaderEXT* pExecutionSetWrites);
```

## [](#_parameters)Parameters

- `device` is the logical device that owns the indirect execution set.
- `indirectExecutionSet` is the indirect execution set being updated.
- `executionSetWriteCount` is the number of elements in the `pExecutionSetWrites` array.
- `pExecutionSetWrites` is a pointer to an array of [VkWriteIndirectExecutionSetShaderEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWriteIndirectExecutionSetShaderEXT.html) structures describing the elements to update.

## [](#_description)Description

Valid Usage

- [](#VUID-vkUpdateIndirectExecutionSetShaderEXT-indirectExecutionSet-11041)VUID-vkUpdateIndirectExecutionSetShaderEXT-indirectExecutionSet-11041  
  `indirectExecutionSet` **must** have been created with type `VK_INDIRECT_EXECUTION_SET_INFO_TYPE_SHADER_OBJECTS_EXT`
- [](#VUID-vkUpdateIndirectExecutionSetShaderEXT-pExecutionSetWrites-11043)VUID-vkUpdateIndirectExecutionSetShaderEXT-pExecutionSetWrites-11043  
  Each element in the `pExecutionSetWrites` array must have a unique `VkWriteIndirectExecutionSetShaderEXT`::`index`
- [](#VUID-vkUpdateIndirectExecutionSetShaderEXT-None-11044)VUID-vkUpdateIndirectExecutionSetShaderEXT-None-11044  
  Each member of the Indirect Execution Set referenced by the update command **must** not be in use by the device
- [](#VUID-vkUpdateIndirectExecutionSetShaderEXT-pExecutionSetWrites-11140)VUID-vkUpdateIndirectExecutionSetShaderEXT-pExecutionSetWrites-11140  
  The descriptor layout of each shader in `pExecutionSetWrites` **must** be [compatible](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptorsets-compatibility) with the initial layout info used to create the Indirect Execution Set
- [](#VUID-vkUpdateIndirectExecutionSetShaderEXT-None-11148)VUID-vkUpdateIndirectExecutionSetShaderEXT-None-11148  
  Each fragment shader element in the Indirect Execution Set **must** have [identically defined](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#glossary-identically-defined) [fragment outputs interface](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-fragmentoutput) to the initial shader state used to create the Indirect Execution Set
- [](#VUID-vkUpdateIndirectExecutionSetShaderEXT-FragDepth-11054)VUID-vkUpdateIndirectExecutionSetShaderEXT-FragDepth-11054  
  Each fragment shader element in the Indirect Execution Set **must** match the initial shader state used to create the Indirect Execution Set in its use of `FragDepth`
- [](#VUID-vkUpdateIndirectExecutionSetShaderEXT-SampleMask-11050)VUID-vkUpdateIndirectExecutionSetShaderEXT-SampleMask-11050  
  Each fragment shader element in the Indirect Execution Set **must** match the initial shader state used to create the Indirect Execution Set in its use of `SampleMask`
- [](#VUID-vkUpdateIndirectExecutionSetShaderEXT-StencilExportEXT-11003)VUID-vkUpdateIndirectExecutionSetShaderEXT-StencilExportEXT-11003  
  Each fragment shader element in the Indirect Execution Set **must** match the initial shader state used to create the Indirect Execution Set in its use of `StencilExportEXT`

Valid Usage (Implicit)

- [](#VUID-vkUpdateIndirectExecutionSetShaderEXT-device-parameter)VUID-vkUpdateIndirectExecutionSetShaderEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkUpdateIndirectExecutionSetShaderEXT-indirectExecutionSet-parameter)VUID-vkUpdateIndirectExecutionSetShaderEXT-indirectExecutionSet-parameter  
  `indirectExecutionSet` **must** be a valid [VkIndirectExecutionSetEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectExecutionSetEXT.html) handle
- [](#VUID-vkUpdateIndirectExecutionSetShaderEXT-pExecutionSetWrites-parameter)VUID-vkUpdateIndirectExecutionSetShaderEXT-pExecutionSetWrites-parameter  
  `pExecutionSetWrites` **must** be a valid pointer to an array of `executionSetWriteCount` valid [VkWriteIndirectExecutionSetShaderEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWriteIndirectExecutionSetShaderEXT.html) structures
- [](#VUID-vkUpdateIndirectExecutionSetShaderEXT-executionSetWriteCount-arraylength)VUID-vkUpdateIndirectExecutionSetShaderEXT-executionSetWriteCount-arraylength  
  `executionSetWriteCount` **must** be greater than `0`
- [](#VUID-vkUpdateIndirectExecutionSetShaderEXT-indirectExecutionSet-parent)VUID-vkUpdateIndirectExecutionSetShaderEXT-indirectExecutionSet-parent  
  `indirectExecutionSet` **must** have been created, allocated, or retrieved from `device`

Host Synchronization

- Host access to `indirectExecutionSet` **must** be externally synchronized

## [](#_see_also)See Also

[VK\_EXT\_device\_generated\_commands](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_device_generated_commands.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkIndirectExecutionSetEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectExecutionSetEXT.html), [VkWriteIndirectExecutionSetShaderEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWriteIndirectExecutionSetShaderEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkUpdateIndirectExecutionSetShaderEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0