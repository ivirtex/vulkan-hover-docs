# VkMutableDescriptorTypeListEXT(3) Manual Page

## Name

VkMutableDescriptorTypeListEXT - Structure describing descriptor types that a given descriptor may mutate to



## [](#_c_specification)C Specification

The list of potential descriptor types a given mutable descriptor **can** mutate to is passed in a `VkMutableDescriptorTypeListEXT` structure.

The `VkMutableDescriptorTypeListEXT` structure is defined as:

```c++
// Provided by VK_EXT_mutable_descriptor_type
typedef struct VkMutableDescriptorTypeListEXT {
    uint32_t                   descriptorTypeCount;
    const VkDescriptorType*    pDescriptorTypes;
} VkMutableDescriptorTypeListEXT;
```

or the equivalent

```c++
// Provided by VK_VALVE_mutable_descriptor_type
typedef VkMutableDescriptorTypeListEXT VkMutableDescriptorTypeListVALVE;
```

## [](#_members)Members

- `descriptorTypeCount` is the number of elements in `pDescriptorTypes`.
- `pDescriptorTypes` is `NULL` or a pointer to an array of `descriptorTypeCount` [VkDescriptorType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorType.html) values defining which descriptor types a given binding may mutate to.

## [](#_description)Description

Valid Usage

- [](#VUID-VkMutableDescriptorTypeListEXT-descriptorTypeCount-04597)VUID-VkMutableDescriptorTypeListEXT-descriptorTypeCount-04597  
  `descriptorTypeCount` **must** not be `0` if the corresponding binding is of `VK_DESCRIPTOR_TYPE_MUTABLE_EXT`
- [](#VUID-VkMutableDescriptorTypeListEXT-pDescriptorTypes-04598)VUID-VkMutableDescriptorTypeListEXT-pDescriptorTypes-04598  
  `pDescriptorTypes` **must** be a valid pointer to an array of `descriptorTypeCount` valid, unique [VkDescriptorType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorType.html) values if the given binding is of `VK_DESCRIPTOR_TYPE_MUTABLE_EXT` type
- [](#VUID-VkMutableDescriptorTypeListEXT-descriptorTypeCount-04599)VUID-VkMutableDescriptorTypeListEXT-descriptorTypeCount-04599  
  `descriptorTypeCount` **must** be `0` if the corresponding binding is not of `VK_DESCRIPTOR_TYPE_MUTABLE_EXT`
- [](#VUID-VkMutableDescriptorTypeListEXT-pDescriptorTypes-04600)VUID-VkMutableDescriptorTypeListEXT-pDescriptorTypes-04600  
  `pDescriptorTypes` **must** not contain `VK_DESCRIPTOR_TYPE_MUTABLE_EXT`
- [](#VUID-VkMutableDescriptorTypeListEXT-pDescriptorTypes-04601)VUID-VkMutableDescriptorTypeListEXT-pDescriptorTypes-04601  
  `pDescriptorTypes` **must** not contain `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER_DYNAMIC`
- [](#VUID-VkMutableDescriptorTypeListEXT-pDescriptorTypes-04602)VUID-VkMutableDescriptorTypeListEXT-pDescriptorTypes-04602  
  `pDescriptorTypes` **must** not contain `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER_DYNAMIC`
- [](#VUID-VkMutableDescriptorTypeListEXT-pDescriptorTypes-04603)VUID-VkMutableDescriptorTypeListEXT-pDescriptorTypes-04603  
  `pDescriptorTypes` **must** not contain `VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK`
- [](#VUID-VkMutableDescriptorTypeListEXT-pDescriptorTypes-09696)VUID-VkMutableDescriptorTypeListEXT-pDescriptorTypes-09696  
  `pDescriptorTypes` **must** not contain `VK_DESCRIPTOR_TYPE_TENSOR_ARM`

Valid Usage (Implicit)

- [](#VUID-VkMutableDescriptorTypeListEXT-pDescriptorTypes-parameter)VUID-VkMutableDescriptorTypeListEXT-pDescriptorTypes-parameter  
  If `descriptorTypeCount` is not `0`, `pDescriptorTypes` **must** be a valid pointer to an array of `descriptorTypeCount` valid [VkDescriptorType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorType.html) values

## [](#_see_also)See Also

[VK\_EXT\_mutable\_descriptor\_type](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_mutable_descriptor_type.html), [VK\_VALVE\_mutable\_descriptor\_type](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VALVE_mutable_descriptor_type.html), [VkDescriptorType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorType.html), [VkMutableDescriptorTypeCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMutableDescriptorTypeCreateInfoEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkMutableDescriptorTypeListEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0