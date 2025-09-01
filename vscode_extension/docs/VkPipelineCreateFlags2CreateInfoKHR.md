# VkPipelineCreateFlags2CreateInfo(3) Manual Page

## Name

VkPipelineCreateFlags2CreateInfo - Extended pipeline create flags



## [](#_c_specification)C Specification

The `VkPipelineCreateFlags2CreateInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_4
typedef struct VkPipelineCreateFlags2CreateInfo {
    VkStructureType           sType;
    const void*               pNext;
    VkPipelineCreateFlags2    flags;
} VkPipelineCreateFlags2CreateInfo;
```

or the equivalent

```c++
// Provided by VK_KHR_maintenance5
typedef VkPipelineCreateFlags2CreateInfo VkPipelineCreateFlags2CreateInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is a bitmask of [VkPipelineCreateFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateFlagBits2.html) specifying how a pipeline will be generated.

## [](#_description)Description

If this structure is included in the `pNext` chain of a pipeline creation structure, `flags` is used instead of the corresponding `flags` value passed in that creation structure, allowing additional creation flags to be specified.

Valid Usage (Implicit)

- [](#VUID-VkPipelineCreateFlags2CreateInfo-sType-sType)VUID-VkPipelineCreateFlags2CreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PIPELINE_CREATE_FLAGS_2_CREATE_INFO`
- [](#VUID-VkPipelineCreateFlags2CreateInfo-flags-parameter)VUID-VkPipelineCreateFlags2CreateInfo-flags-parameter  
  `flags` **must** be a valid combination of [VkPipelineCreateFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateFlagBits2.html) values
- [](#VUID-VkPipelineCreateFlags2CreateInfo-flags-requiredbitmask)VUID-VkPipelineCreateFlags2CreateInfo-flags-requiredbitmask  
  `flags` **must** not be `0`

## [](#_see_also)See Also

[VK\_KHR\_maintenance5](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance5.html), [VK\_VERSION\_1\_4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_4.html), [VkPipelineCreateFlags2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateFlags2.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPipelineCreateFlags2CreateInfo).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0