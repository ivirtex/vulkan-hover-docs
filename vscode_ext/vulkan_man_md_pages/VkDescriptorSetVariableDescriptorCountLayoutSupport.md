# VkDescriptorSetVariableDescriptorCountLayoutSupport(3) Manual Page

## Name

VkDescriptorSetVariableDescriptorCountLayoutSupport - Structure
returning information about whether a descriptor set layout can be
supported



## <a href="#_c_specification" class="anchor"></a>C Specification

If the `pNext` chain of a
[VkDescriptorSetLayoutSupport](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayoutSupport.html)
structure includes a
`VkDescriptorSetVariableDescriptorCountLayoutSupport` structure, then
that structure returns additional information about whether the
descriptor set layout is supported.

``` c
// Provided by VK_VERSION_1_2
typedef struct VkDescriptorSetVariableDescriptorCountLayoutSupport {
    VkStructureType    sType;
    void*              pNext;
    uint32_t           maxVariableDescriptorCount;
} VkDescriptorSetVariableDescriptorCountLayoutSupport;
```

or the equivalent

``` c
// Provided by VK_EXT_descriptor_indexing
typedef VkDescriptorSetVariableDescriptorCountLayoutSupport VkDescriptorSetVariableDescriptorCountLayoutSupportEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `maxVariableDescriptorCount` indicates the maximum number of
  descriptors supported in the highest numbered binding of the layout,
  if that binding is variable-sized. If the highest numbered binding of
  the layout has a descriptor type of
  `VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK` then
  `maxVariableDescriptorCount` indicates the maximum byte size supported
  for the binding, if that binding is variable-sized.

## <a href="#_description" class="anchor"></a>Description

If the
[VkDescriptorSetLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayoutCreateInfo.html)
structure specified in
[vkGetDescriptorSetLayoutSupport](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDescriptorSetLayoutSupport.html)::`pCreateInfo`
includes a variable-sized descriptor, then `supported` is determined
assuming the requested size of the variable-sized descriptor, and
`maxVariableDescriptorCount` is set to the maximum size of that
descriptor that **can** be successfully created (which is greater than
or equal to the requested size passed in). If the
[VkDescriptorSetLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayoutCreateInfo.html)
structure does not include a variable-sized descriptor, or if the
[VkPhysicalDeviceDescriptorIndexingFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDescriptorIndexingFeatures.html)::`descriptorBindingVariableDescriptorCount`
feature is not enabled, then `maxVariableDescriptorCount` is set to
zero. For the purposes of this command, a variable-sized descriptor
binding with a `descriptorCount` of zero is treated as having a
`descriptorCount` of four if `descriptorType` is
`VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK`, or one otherwise, and thus
the binding is not ignored and the maximum descriptor count will be
returned. If the layout is not supported, then the value written to
`maxVariableDescriptorCount` is undefined.

Valid Usage (Implicit)

- <a
  href="#VUID-VkDescriptorSetVariableDescriptorCountLayoutSupport-sType-sType"
  id="VUID-VkDescriptorSetVariableDescriptorCountLayoutSupport-sType-sType"></a>
  VUID-VkDescriptorSetVariableDescriptorCountLayoutSupport-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_DESCRIPTOR_SET_VARIABLE_DESCRIPTOR_COUNT_LAYOUT_SUPPORT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_descriptor_indexing](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_descriptor_indexing.html),
[VK_VERSION_1_2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_2.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDescriptorSetVariableDescriptorCountLayoutSupport"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
