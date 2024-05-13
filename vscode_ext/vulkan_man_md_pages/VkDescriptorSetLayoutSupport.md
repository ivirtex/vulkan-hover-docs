# VkDescriptorSetLayoutSupport(3) Manual Page

## Name

VkDescriptorSetLayoutSupport - Structure returning information about
whether a descriptor set layout can be supported



## <a href="#_c_specification" class="anchor"></a>C Specification

Information about support for the descriptor set layout is returned in a
`VkDescriptorSetLayoutSupport` structure:

``` c
// Provided by VK_VERSION_1_1
typedef struct VkDescriptorSetLayoutSupport {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           supported;
} VkDescriptorSetLayoutSupport;
```

or the equivalent

``` c
// Provided by VK_KHR_maintenance3
typedef VkDescriptorSetLayoutSupport VkDescriptorSetLayoutSupportKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `supported` specifies whether the descriptor set layout **can** be
  created.

## <a href="#_description" class="anchor"></a>Description

`supported` is set to `VK_TRUE` if the descriptor set **can** be
created, or else is set to `VK_FALSE`.

Valid Usage (Implicit)

- <a href="#VUID-VkDescriptorSetLayoutSupport-sType-sType"
  id="VUID-VkDescriptorSetLayoutSupport-sType-sType"></a>
  VUID-VkDescriptorSetLayoutSupport-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DESCRIPTOR_SET_LAYOUT_SUPPORT`

- <a href="#VUID-VkDescriptorSetLayoutSupport-pNext-pNext"
  id="VUID-VkDescriptorSetLayoutSupport-pNext-pNext"></a>
  VUID-VkDescriptorSetLayoutSupport-pNext-pNext  
  `pNext` **must** be `NULL` or a pointer to a valid instance of
  [VkDescriptorSetVariableDescriptorCountLayoutSupport](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetVariableDescriptorCountLayoutSupport.html)

- <a href="#VUID-VkDescriptorSetLayoutSupport-sType-unique"
  id="VUID-VkDescriptorSetLayoutSupport-sType-unique"></a>
  VUID-VkDescriptorSetLayoutSupport-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html), [VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetDescriptorSetLayoutSupport](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDescriptorSetLayoutSupport.html),
[vkGetDescriptorSetLayoutSupportKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDescriptorSetLayoutSupportKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDescriptorSetLayoutSupport"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
