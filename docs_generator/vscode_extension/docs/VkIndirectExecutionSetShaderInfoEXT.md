# VkIndirectExecutionSetShaderInfoEXT(3) Manual Page

## Name

VkIndirectExecutionSetShaderInfoEXT - Struct specifying parameters of a newly created indirect execution set containing only shader objects



## [](#_c_specification)C Specification

The `VkIndirectExecutionSetShaderInfoEXT` structure is defined as:

```c++
// Provided by VK_EXT_device_generated_commands
typedef struct VkIndirectExecutionSetShaderInfoEXT {
    VkStructureType                                     sType;
    const void*                                         pNext;
    uint32_t                                            shaderCount;
    const VkShaderEXT*                                  pInitialShaders;
    const VkIndirectExecutionSetShaderLayoutInfoEXT*    pSetLayoutInfos;
    uint32_t                                            maxShaderCount;
    uint32_t                                            pushConstantRangeCount;
    const VkPushConstantRange*                          pPushConstantRanges;
} VkIndirectExecutionSetShaderInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `shaderCount` is the number of members in the `pInitialShaders` and `pSetLayoutInfos` arrays.
- `pInitialShaders` is a pointer to an array containing a [VkShaderEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderEXT.html) object for each shader stage that will be used in the set. These shaders will be automatically added to the set beginning at index `0`.
- `pSetLayoutInfos` is a pointer to an array containing a [VkIndirectExecutionSetShaderLayoutInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectExecutionSetShaderLayoutInfoEXT.html) used by each corresponding `pInitialShaders` shader stage in the set.
- `maxShaderCount` is the maximum number of shader objects stored in the set.
- `pushConstantRangeCount` is the number of members in the `pPushConstantRanges` array.
- `pPushConstantRanges` is a pointer to the array of [VkPushConstantRange](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPushConstantRange.html) ranges used by all shaders in the set.

## [](#_description)Description

The characteristics of `pInitialShaders` will be used to validate all shaders added to the set even if they are removed from the set or destroyed.

When an Indirect Execution Set created with shader objects is used, `pInitialShaders` constitutes the initial shader state.

Valid Usage

- [](#VUID-VkIndirectExecutionSetShaderInfoEXT-pInitialShaders-11020)VUID-VkIndirectExecutionSetShaderInfoEXT-pInitialShaders-11020  
  All members of `pInitialShaders` **must** have a `stage` supported by [](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-supportedIndirectCommandsShaderStagesShaderBinding)[VkPhysicalDeviceDeviceGeneratedCommandsPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDeviceGeneratedCommandsPropertiesEXT.html)::`supportedIndirectCommandsShaderStagesShaderBinding`
- [](#VUID-VkIndirectExecutionSetShaderInfoEXT-maxShaderCount-11021)VUID-VkIndirectExecutionSetShaderInfoEXT-maxShaderCount-11021  
  `maxShaderCount` **must** not be zero
- [](#VUID-VkIndirectExecutionSetShaderInfoEXT-maxShaderCount-11022)VUID-VkIndirectExecutionSetShaderInfoEXT-maxShaderCount-11022  
  `maxShaderCount` **must** be less than or equal to `VkPhysicalDeviceDeviceGeneratedCommandsPropertiesEXT`::`maxIndirectShaderObjectCount`
- [](#VUID-VkIndirectExecutionSetShaderInfoEXT-maxShaderCount-11036)VUID-VkIndirectExecutionSetShaderInfoEXT-maxShaderCount-11036  
  `maxShaderCount` **must** be greater than or equal to `shaderCount`
- [](#VUID-VkIndirectExecutionSetShaderInfoEXT-stage-11023)VUID-VkIndirectExecutionSetShaderInfoEXT-stage-11023  
  The `stage` of each element in the `pInitialShaders` array **must** be unique
- [](#VUID-VkIndirectExecutionSetShaderInfoEXT-pInitialShaders-11154)VUID-VkIndirectExecutionSetShaderInfoEXT-pInitialShaders-11154  
  Each member of `pInitialShaders` **must** have been created with `VK_SHADER_CREATE_INDIRECT_BINDABLE_BIT_EXT`

Valid Usage (Implicit)

- [](#VUID-VkIndirectExecutionSetShaderInfoEXT-sType-sType)VUID-VkIndirectExecutionSetShaderInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_INDIRECT_EXECUTION_SET_SHADER_INFO_EXT`
- [](#VUID-VkIndirectExecutionSetShaderInfoEXT-pInitialShaders-parameter)VUID-VkIndirectExecutionSetShaderInfoEXT-pInitialShaders-parameter  
  `pInitialShaders` **must** be a valid pointer to an array of `shaderCount` valid [VkShaderEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderEXT.html) handles
- [](#VUID-VkIndirectExecutionSetShaderInfoEXT-pSetLayoutInfos-parameter)VUID-VkIndirectExecutionSetShaderInfoEXT-pSetLayoutInfos-parameter  
  If `pSetLayoutInfos` is not `NULL`, `pSetLayoutInfos` **must** be a valid pointer to an array of `shaderCount` valid [VkIndirectExecutionSetShaderLayoutInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectExecutionSetShaderLayoutInfoEXT.html) structures
- [](#VUID-VkIndirectExecutionSetShaderInfoEXT-pPushConstantRanges-parameter)VUID-VkIndirectExecutionSetShaderInfoEXT-pPushConstantRanges-parameter  
  If `pushConstantRangeCount` is not `0`, `pPushConstantRanges` **must** be a valid pointer to an array of `pushConstantRangeCount` valid [VkPushConstantRange](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPushConstantRange.html) structures
- [](#VUID-VkIndirectExecutionSetShaderInfoEXT-shaderCount-arraylength)VUID-VkIndirectExecutionSetShaderInfoEXT-shaderCount-arraylength  
  `shaderCount` **must** be greater than `0`

## [](#_see_also)See Also

[VK\_EXT\_device\_generated\_commands](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_device_generated_commands.html), [VkIndirectExecutionSetInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectExecutionSetInfoEXT.html), [VkIndirectExecutionSetShaderLayoutInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectExecutionSetShaderLayoutInfoEXT.html), [VkPushConstantRange](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPushConstantRange.html), [VkShaderEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderEXT.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkIndirectExecutionSetShaderInfoEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0