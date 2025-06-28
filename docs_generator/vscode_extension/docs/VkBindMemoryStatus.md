# VkBindMemoryStatus(3) Manual Page

## Name

VkBindMemoryStatus - Structure specifying where to return memory binding status



## [](#_c_specification)C Specification

The `VkBindMemoryStatus` structure is defined as:

```c++
// Provided by VK_VERSION_1_4
typedef struct VkBindMemoryStatus {
    VkStructureType    sType;
    const void*        pNext;
    VkResult*          pResult;
} VkBindMemoryStatus;
```

or the equivalent

```c++
// Provided by VK_KHR_maintenance6
typedef VkBindMemoryStatus VkBindMemoryStatusKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `pResult` is a pointer to a `VkResult` value.

## [](#_description)Description

If the `pNext` chain of [VkBindBufferMemoryInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindBufferMemoryInfo.html) or [VkBindImageMemoryInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindImageMemoryInfo.html) includes a `VkBindMemoryStatus` structure, then the `VkBindMemoryStatus`::`pResult` will be populated with a value describing the result of the corresponding memory binding operation.

Valid Usage (Implicit)

- [](#VUID-VkBindMemoryStatus-sType-sType)VUID-VkBindMemoryStatus-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_BIND_MEMORY_STATUS`
- [](#VUID-VkBindMemoryStatus-pResult-parameter)VUID-VkBindMemoryStatus-pResult-parameter  
  `pResult` **must** be a valid pointer to a [VkResult](https://registry.khronos.org/vulkan/specs/latest/man/html/VkResult.html) value

## [](#_see_also)See Also

[VK\_KHR\_maintenance6](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance6.html), [VK\_VERSION\_1\_4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_4.html), [VkResult](https://registry.khronos.org/vulkan/specs/latest/man/html/VkResult.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkBindMemoryStatus)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0