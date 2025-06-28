# VkPhysicalDeviceDeviceGeneratedCommandsPropertiesNV(3) Manual Page

## Name

VkPhysicalDeviceDeviceGeneratedCommandsPropertiesNV - Structure describing push descriptor limits that can be supported by an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceDeviceGeneratedCommandsPropertiesNV` structure is defined as:

```c++
// Provided by VK_NV_device_generated_commands
typedef struct VkPhysicalDeviceDeviceGeneratedCommandsPropertiesNV {
    VkStructureType    sType;
    void*              pNext;
    uint32_t           maxGraphicsShaderGroupCount;
    uint32_t           maxIndirectSequenceCount;
    uint32_t           maxIndirectCommandsTokenCount;
    uint32_t           maxIndirectCommandsStreamCount;
    uint32_t           maxIndirectCommandsTokenOffset;
    uint32_t           maxIndirectCommandsStreamStride;
    uint32_t           minSequencesCountBufferOffsetAlignment;
    uint32_t           minSequencesIndexBufferOffsetAlignment;
    uint32_t           minIndirectCommandsBufferOffsetAlignment;
} VkPhysicalDeviceDeviceGeneratedCommandsPropertiesNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `maxGraphicsShaderGroupCount` is the maximum number of shader groups in [VkGraphicsPipelineShaderGroupsCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineShaderGroupsCreateInfoNV.html).
- `maxIndirectSequenceCount` is the maximum number of sequences in [VkGeneratedCommandsInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeneratedCommandsInfoNV.html) and in [VkGeneratedCommandsMemoryRequirementsInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeneratedCommandsMemoryRequirementsInfoNV.html).
- `maxIndirectCommandsTokenCount` is the maximum number of tokens in [VkIndirectCommandsLayoutCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsLayoutCreateInfoNV.html).
- `maxIndirectCommandsStreamCount` is the maximum number of streams in [VkIndirectCommandsLayoutCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsLayoutCreateInfoNV.html).
- `maxIndirectCommandsTokenOffset` is the maximum offset in [VkIndirectCommandsLayoutTokenNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsLayoutTokenNV.html).
- `maxIndirectCommandsStreamStride` is the maximum stream stride in [VkIndirectCommandsLayoutCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsLayoutCreateInfoNV.html).
- `minSequencesCountBufferOffsetAlignment` is the minimum alignment for memory addresses which **can** be used in [VkGeneratedCommandsInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeneratedCommandsInfoNV.html).
- `minSequencesIndexBufferOffsetAlignment` is the minimum alignment for memory addresses which **can** be used in [VkGeneratedCommandsInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeneratedCommandsInfoNV.html).
- `minIndirectCommandsBufferOffsetAlignment` is the minimum alignment for memory addresses used in [VkIndirectCommandsStreamNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsStreamNV.html), and as preprocess buffer in [VkGeneratedCommandsInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeneratedCommandsInfoNV.html).

## [](#_description)Description

If the `VkPhysicalDeviceDeviceGeneratedCommandsPropertiesNV` structure is included in the `pNext` chain of the [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html) structure passed to [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2.html), it is filled in with each corresponding implementation-dependent property.

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceDeviceGeneratedCommandsPropertiesNV-sType-sType)VUID-VkPhysicalDeviceDeviceGeneratedCommandsPropertiesNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DEVICE_GENERATED_COMMANDS_PROPERTIES_NV`

## [](#_see_also)See Also

[VK\_NV\_device\_generated\_commands](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_device_generated_commands.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceDeviceGeneratedCommandsPropertiesNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0