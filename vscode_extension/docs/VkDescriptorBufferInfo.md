# VkDescriptorBufferInfo(3) Manual Page

## Name

VkDescriptorBufferInfo - Structure specifying descriptor buffer information



## [](#_c_specification)C Specification

The `VkDescriptorBufferInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_0
typedef struct VkDescriptorBufferInfo {
    VkBuffer        buffer;
    VkDeviceSize    offset;
    VkDeviceSize    range;
} VkDescriptorBufferInfo;
```

## [](#_members)Members

- `buffer` is [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) or the buffer resource.
- `offset` is the offset in bytes from the start of `buffer`. Access to buffer memory via this descriptor uses addressing that is relative to this starting offset.
- `range` is the size in bytes that is used for this descriptor update, or `VK_WHOLE_SIZE` to use the range from `offset` to the end of the buffer.
  
  Note
  
  When setting `range` to `VK_WHOLE_SIZE`, the [effective range](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#buffer-info-effective-range) **must** not be larger than the maximum range for the descriptor type ([`maxUniformBufferRange`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxUniformBufferRange) or [`maxStorageBufferRange`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxStorageBufferRange)). This means that `VK_WHOLE_SIZE` is not typically useful in the common case where uniform buffer descriptors are suballocated from a buffer that is much larger than `maxUniformBufferRange`.

## [](#_description)Description

For `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER_DYNAMIC` and `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER_DYNAMIC` descriptor types, `offset` is the base offset from which the dynamic offset is applied and `range` is the static size used for all dynamic offsets.

When `range` is `VK_WHOLE_SIZE` the effective range is calculated at [vkUpdateDescriptorSets](https://registry.khronos.org/vulkan/specs/latest/man/html/vkUpdateDescriptorSets.html) is by taking the size of `buffer` minus the `offset`.

Valid Usage

- [](#VUID-VkDescriptorBufferInfo-offset-00340)VUID-VkDescriptorBufferInfo-offset-00340  
  `offset` **must** be less than the size of `buffer`
- [](#VUID-VkDescriptorBufferInfo-range-00341)VUID-VkDescriptorBufferInfo-range-00341  
  If `range` is not equal to `VK_WHOLE_SIZE`, `range` **must** be greater than `0`
- [](#VUID-VkDescriptorBufferInfo-range-00342)VUID-VkDescriptorBufferInfo-range-00342  
  If `range` is not equal to `VK_WHOLE_SIZE`, `range` **must** be less than or equal to the size of `buffer` minus `offset`
- [](#VUID-VkDescriptorBufferInfo-buffer-02998)VUID-VkDescriptorBufferInfo-buffer-02998  
  If the [`nullDescriptor`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-nullDescriptor) feature is not enabled, `buffer` **must** not be [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html)
- [](#VUID-VkDescriptorBufferInfo-buffer-02999)VUID-VkDescriptorBufferInfo-buffer-02999  
  If `buffer` is [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `offset` **must** be zero and `range` **must** be `VK_WHOLE_SIZE`

Valid Usage (Implicit)

- [](#VUID-VkDescriptorBufferInfo-buffer-parameter)VUID-VkDescriptorBufferInfo-buffer-parameter  
  If `buffer` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `buffer` **must** be a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) handle

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html), [VkWriteDescriptorSet](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWriteDescriptorSet.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDescriptorBufferInfo).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0