# VkBufferUsageFlags2CreateInfo(3) Manual Page

## Name

VkBufferUsageFlags2CreateInfo - Extended buffer usage flags



## [](#_c_specification)C Specification

The `VkBufferUsageFlags2CreateInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_4
typedef struct VkBufferUsageFlags2CreateInfo {
    VkStructureType        sType;
    const void*            pNext;
    VkBufferUsageFlags2    usage;
} VkBufferUsageFlags2CreateInfo;
```

or the equivalent

```c++
// Provided by VK_KHR_maintenance5
typedef VkBufferUsageFlags2CreateInfo VkBufferUsageFlags2CreateInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `usage` is a bitmask of [VkBufferUsageFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferUsageFlagBits2.html) specifying allowed usages of the buffer.

## [](#_description)Description

If this structure is included in the `pNext` chain of a buffer creation structure, `usage` is used instead of the corresponding `usage` value passed in that creation structure, allowing additional usage flags to be specified. If this structure is included in the `pNext` chain of a buffer query structure, the usage flags of the buffer are returned in `usage` of this structure, and the usage flags representable in `usage` of the buffer query structure are also returned in that field.

Valid Usage (Implicit)

- [](#VUID-VkBufferUsageFlags2CreateInfo-sType-sType)VUID-VkBufferUsageFlags2CreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_BUFFER_USAGE_FLAGS_2_CREATE_INFO`
- [](#VUID-VkBufferUsageFlags2CreateInfo-usage-parameter)VUID-VkBufferUsageFlags2CreateInfo-usage-parameter  
  `usage` **must** be a valid combination of [VkBufferUsageFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferUsageFlagBits2.html) values
- [](#VUID-VkBufferUsageFlags2CreateInfo-usage-requiredbitmask)VUID-VkBufferUsageFlags2CreateInfo-usage-requiredbitmask  
  `usage` **must** not be `0`

## [](#_see_also)See Also

[VK\_KHR\_maintenance5](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance5.html), [VK\_VERSION\_1\_4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_4.html), [VkBufferUsageFlags2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferUsageFlags2.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkBufferUsageFlags2CreateInfo)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0