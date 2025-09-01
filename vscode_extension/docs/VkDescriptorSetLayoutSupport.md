# VkDescriptorSetLayoutSupport(3) Manual Page

## Name

VkDescriptorSetLayoutSupport - Structure returning information about whether a descriptor set layout can be supported



## [](#_c_specification)C Specification

Information about support for the descriptor set layout is returned in a `VkDescriptorSetLayoutSupport` structure:

```c++
// Provided by VK_VERSION_1_1
typedef struct VkDescriptorSetLayoutSupport {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           supported;
} VkDescriptorSetLayoutSupport;
```

or the equivalent

```c++
// Provided by VK_KHR_maintenance3
typedef VkDescriptorSetLayoutSupport VkDescriptorSetLayoutSupportKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `supported` specifies whether the descriptor set layout **can** be created.

## [](#_description)Description

`supported` will be `VK_TRUE` if the descriptor set **can** be created, or else `VK_FALSE`.

Valid Usage (Implicit)

- [](#VUID-VkDescriptorSetLayoutSupport-sType-sType)VUID-VkDescriptorSetLayoutSupport-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_SUPPORT`
- [](#VUID-VkDescriptorSetLayoutSupport-pNext-pNext)VUID-VkDescriptorSetLayoutSupport-pNext-pNext  
  `pNext` **must** be `NULL` or a pointer to a valid instance of [VkDescriptorSetVariableDescriptorCountLayoutSupport](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetVariableDescriptorCountLayoutSupport.html)
- [](#VUID-VkDescriptorSetLayoutSupport-sType-unique)VUID-VkDescriptorSetLayoutSupport-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique

## [](#_see_also)See Also

[VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetDescriptorSetLayoutSupport](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDescriptorSetLayoutSupport.html), [vkGetDescriptorSetLayoutSupport](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDescriptorSetLayoutSupport.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDescriptorSetLayoutSupport).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0