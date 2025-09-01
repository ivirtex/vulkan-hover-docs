# VkDescriptorUpdateTemplateEntry(3) Manual Page

## Name

VkDescriptorUpdateTemplateEntry - Describes a single descriptor update of the descriptor update template



## [](#_c_specification)C Specification

The `VkDescriptorUpdateTemplateEntry` structure is defined as:

```c++
// Provided by VK_VERSION_1_1
typedef struct VkDescriptorUpdateTemplateEntry {
    uint32_t            dstBinding;
    uint32_t            dstArrayElement;
    uint32_t            descriptorCount;
    VkDescriptorType    descriptorType;
    size_t              offset;
    size_t              stride;
} VkDescriptorUpdateTemplateEntry;
```

or the equivalent

```c++
// Provided by VK_KHR_descriptor_update_template
typedef VkDescriptorUpdateTemplateEntry VkDescriptorUpdateTemplateEntryKHR;
```

## [](#_members)Members

- `dstBinding` is the descriptor binding to update when using this descriptor update template.
- `dstArrayElement` is the starting element in the array belonging to `dstBinding`. If the descriptor binding identified by `dstBinding` has a descriptor type of `VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK` then `dstArrayElement` specifies the starting byte offset to update.
- `descriptorCount` is the number of descriptors to update. If `descriptorCount` is greater than the number of remaining array elements in the destination binding, those affect consecutive bindings in a manner similar to [VkWriteDescriptorSet](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWriteDescriptorSet.html) above. If the descriptor binding identified by `dstBinding` has a descriptor type of `VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK` then `descriptorCount` specifies the number of bytes to update and the remaining array elements in the destination binding refer to the remaining number of bytes in it.
- `descriptorType` is a [VkDescriptorType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorType.html) specifying the type of the descriptor.
- `offset` is the offset in bytes of the first binding in the raw data structure.
- `stride` is the stride in bytes between two consecutive array elements of the descriptor update information in the raw data structure. The actual pointer ptr for each array element j of update entry i is computed using the following formula:
  
  ```c++
      const char *ptr = (const char *)pData + pDescriptorUpdateEntries[i].offset + j * pDescriptorUpdateEntries[i].stride
  ```
  
  The stride is useful in case the bindings are stored in structs along with other data. If `descriptorType` is `VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK` then the value of `stride` is ignored and the stride is assumed to be `1`, i.e. the descriptor update information for them is always specified as a contiguous range.

## [](#_description)Description

Valid Usage

- [](#VUID-VkDescriptorUpdateTemplateEntry-dstBinding-00354)VUID-VkDescriptorUpdateTemplateEntry-dstBinding-00354  
  `dstBinding` **must** be a valid binding in the descriptor set layout implicitly specified when using a descriptor update template to update descriptors
- [](#VUID-VkDescriptorUpdateTemplateEntry-dstArrayElement-00355)VUID-VkDescriptorUpdateTemplateEntry-dstArrayElement-00355  
  `dstArrayElement` and `descriptorCount` **must** be less than or equal to the number of array elements in the descriptor set binding implicitly specified when using a descriptor update template to update descriptors, and all applicable [consecutive bindings](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptorsets-updates-consecutive)
- [](#VUID-VkDescriptorUpdateTemplateEntry-descriptor-02226)VUID-VkDescriptorUpdateTemplateEntry-descriptor-02226  
  If `descriptor` type is `VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK`, `dstArrayElement` **must** be an integer multiple of `4`
- [](#VUID-VkDescriptorUpdateTemplateEntry-descriptor-02227)VUID-VkDescriptorUpdateTemplateEntry-descriptor-02227  
  If `descriptor` type is `VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK`, `descriptorCount` **must** be an integer multiple of `4`

Valid Usage (Implicit)

- [](#VUID-VkDescriptorUpdateTemplateEntry-descriptorType-parameter)VUID-VkDescriptorUpdateTemplateEntry-descriptorType-parameter  
  `descriptorType` **must** be a valid [VkDescriptorType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorType.html) value

## [](#_see_also)See Also

[VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VkDescriptorType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorType.html), [VkDescriptorUpdateTemplateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorUpdateTemplateCreateInfo.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDescriptorUpdateTemplateEntry).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0