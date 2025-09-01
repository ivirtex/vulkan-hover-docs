# VkPhysicalDeviceDeviceGeneratedCommandsPropertiesEXT(3) Manual Page

## Name

VkPhysicalDeviceDeviceGeneratedCommandsPropertiesEXT - Structure describing push descriptor limits that can be supported by an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceDeviceGeneratedCommandsPropertiesEXT` structure is defined as:

```c++
// Provided by VK_EXT_device_generated_commands
typedef struct VkPhysicalDeviceDeviceGeneratedCommandsPropertiesEXT {
    VkStructureType                        sType;
    void*                                  pNext;
    uint32_t                               maxIndirectPipelineCount;
    uint32_t                               maxIndirectShaderObjectCount;
    uint32_t                               maxIndirectSequenceCount;
    uint32_t                               maxIndirectCommandsTokenCount;
    uint32_t                               maxIndirectCommandsTokenOffset;
    uint32_t                               maxIndirectCommandsIndirectStride;
    VkIndirectCommandsInputModeFlagsEXT    supportedIndirectCommandsInputModes;
    VkShaderStageFlags                     supportedIndirectCommandsShaderStages;
    VkShaderStageFlags                     supportedIndirectCommandsShaderStagesPipelineBinding;
    VkShaderStageFlags                     supportedIndirectCommandsShaderStagesShaderBinding;
    VkBool32                               deviceGeneratedCommandsTransformFeedback;
    VkBool32                               deviceGeneratedCommandsMultiDrawIndirectCount;
} VkPhysicalDeviceDeviceGeneratedCommandsPropertiesEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`maxIndirectPipelineCount` is the maximum number of pipelines passed to [vkCreateIndirectExecutionSetEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateIndirectExecutionSetEXT.html).
- []()`maxIndirectShaderObjectCount` is the maximum number of shader objects passed to [vkCreateIndirectExecutionSetEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateIndirectExecutionSetEXT.html). If this value is zero, binding shader objects indirectly is not supported.
- []()`maxIndirectSequenceCount` is the maximum number of sequences in [VkGeneratedCommandsInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeneratedCommandsInfoEXT.html) and in [VkGeneratedCommandsMemoryRequirementsInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeneratedCommandsMemoryRequirementsInfoEXT.html).
- []()`maxIndirectCommandsTokenCount` is the maximum number of tokens in [VkIndirectCommandsLayoutCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsLayoutCreateInfoEXT.html).
- []()`maxIndirectCommandsTokenOffset` is the maximum offset in [VkIndirectCommandsLayoutTokenEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsLayoutTokenEXT.html).
- []()`maxIndirectCommandsIndirectStride` is the maximum stream stride in [VkIndirectCommandsLayoutCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsLayoutCreateInfoEXT.html).
- []()`supportedIndirectCommandsInputModes` indicates the supported input modes.
- []()`supportedIndirectCommandsShaderStages` indicates the stages which **can** be used to generate indirect commands. Implementations are required to support, at minimum: `VK_SHADER_STAGE_VERTEX_BIT`, `VK_SHADER_STAGE_FRAGMENT_BIT`, `VK_SHADER_STAGE_COMPUTE_BIT`.
- []()`supportedIndirectCommandsShaderStagesPipelineBinding` indicates the stages which **can** be used within indirect execution sets for indirectly binding shader stages using pipelines.
- []()`supportedIndirectCommandsShaderStagesShaderBinding` indicates the stages which **can** be used within indirect execution sets for indirectly binding shader stages using shader objects.
- []()`deviceGeneratedCommandsTransformFeedback` indicates whether the implementation supports interactions with `VK_EXT_transform_feedback` for pipelines not created with `VK_PIPELINE_CREATE_2_INDIRECT_BINDABLE_BIT_EXT`.
- []()`deviceGeneratedCommandsMultiDrawIndirectCount` indicates whether the implementation supports COUNT variants of multi-draw indirect tokens.

## [](#_description)Description

If the `VkPhysicalDeviceDeviceGeneratedCommandsPropertiesEXT` structure is included in the `pNext` chain of the [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html) structure passed to [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2.html), it is filled in with each corresponding implementation-dependent property.

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceDeviceGeneratedCommandsPropertiesEXT-sType-sType)VUID-VkPhysicalDeviceDeviceGeneratedCommandsPropertiesEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DEVICE_GENERATED_COMMANDS_PROPERTIES_EXT`

## [](#_see_also)See Also

[VK\_EXT\_device\_generated\_commands](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_device_generated_commands.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkIndirectCommandsInputModeFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsInputModeFlagsEXT.html), [VkShaderStageFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderStageFlags.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceDeviceGeneratedCommandsPropertiesEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0