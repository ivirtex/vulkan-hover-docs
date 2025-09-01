# VkSemaphoreTypeCreateInfo(3) Manual Page

## Name

VkSemaphoreTypeCreateInfo - Structure specifying the type of a newly created semaphore



## [](#_c_specification)C Specification

The `VkSemaphoreTypeCreateInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_2
typedef struct VkSemaphoreTypeCreateInfo {
    VkStructureType    sType;
    const void*        pNext;
    VkSemaphoreType    semaphoreType;
    uint64_t           initialValue;
} VkSemaphoreTypeCreateInfo;
```

or the equivalent

```c++
// Provided by VK_KHR_timeline_semaphore
typedef VkSemaphoreTypeCreateInfo VkSemaphoreTypeCreateInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `semaphoreType` is a [VkSemaphoreType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreType.html) value specifying the type of the semaphore.
- `initialValue` is the initial payload value if `semaphoreType` is `VK_SEMAPHORE_TYPE_TIMELINE`.

## [](#_description)Description

To create a semaphore of a specific type, add a `VkSemaphoreTypeCreateInfo` structure to the [VkSemaphoreCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreCreateInfo.html)::`pNext` chain.

If no `VkSemaphoreTypeCreateInfo` structure is included in the `pNext` chain of [VkSemaphoreCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreCreateInfo.html), then the created semaphore will have a default [VkSemaphoreType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreType.html) of `VK_SEMAPHORE_TYPE_BINARY`.

Valid Usage

- [](#VUID-VkSemaphoreTypeCreateInfo-timelineSemaphore-03252)VUID-VkSemaphoreTypeCreateInfo-timelineSemaphore-03252  
  If the [`timelineSemaphore`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-timelineSemaphore) feature is not enabled, `semaphoreType` **must** not equal `VK_SEMAPHORE_TYPE_TIMELINE`
- [](#VUID-VkSemaphoreTypeCreateInfo-semaphoreType-03279)VUID-VkSemaphoreTypeCreateInfo-semaphoreType-03279  
  If `semaphoreType` is `VK_SEMAPHORE_TYPE_BINARY`, `initialValue` **must** be zero

Valid Usage (Implicit)

- [](#VUID-VkSemaphoreTypeCreateInfo-sType-sType)VUID-VkSemaphoreTypeCreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SEMAPHORE_TYPE_CREATE_INFO`
- [](#VUID-VkSemaphoreTypeCreateInfo-semaphoreType-parameter)VUID-VkSemaphoreTypeCreateInfo-semaphoreType-parameter  
  `semaphoreType` **must** be a valid [VkSemaphoreType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreType.html) value

## [](#_see_also)See Also

[VK\_KHR\_timeline\_semaphore](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_timeline_semaphore.html), [VK\_VERSION\_1\_2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_2.html), [VkSemaphoreType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreType.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSemaphoreTypeCreateInfo).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0