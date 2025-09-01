# VkExternalMemoryImageCreateInfo(3) Manual Page

## Name

VkExternalMemoryImageCreateInfo - Specify that an image may be backed by external memory



## [](#_c_specification)C Specification

To define a set of external memory handle types that **may** be used as backing store for an image, add a [VkExternalMemoryImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryImageCreateInfo.html) structure to the `pNext` chain of the [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) structure. The `VkExternalMemoryImageCreateInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_1
typedef struct VkExternalMemoryImageCreateInfo {
    VkStructureType                    sType;
    const void*                        pNext;
    VkExternalMemoryHandleTypeFlags    handleTypes;
} VkExternalMemoryImageCreateInfo;
```

or the equivalent

```c++
// Provided by VK_KHR_external_memory
typedef VkExternalMemoryImageCreateInfo VkExternalMemoryImageCreateInfoKHR;
```

## [](#_members)Members

Note

A `VkExternalMemoryImageCreateInfo` structure with a non-zero `handleTypes` field must be included in the creation parameters for an image that will be bound to memory that is either exported or imported.

## [](#_description)Description

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `handleTypes` is zero or a bitmask of [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBits.html) specifying one or more external memory handle types.

Valid Usage (Implicit)

- [](#VUID-VkExternalMemoryImageCreateInfo-sType-sType)VUID-VkExternalMemoryImageCreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_EXTERNAL_MEMORY_IMAGE_CREATE_INFO`
- [](#VUID-VkExternalMemoryImageCreateInfo-handleTypes-parameter)VUID-VkExternalMemoryImageCreateInfo-handleTypes-parameter  
  `handleTypes` **must** be a valid combination of [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBits.html) values

## [](#_see_also)See Also

[VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VkExternalMemoryHandleTypeFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlags.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkExternalMemoryImageCreateInfo).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0