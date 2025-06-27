# VkMutableDescriptorTypeListEXT(3) Manual Page

## Name

VkMutableDescriptorTypeListEXT - Structure describing descriptor types
that a given descriptor may mutate to



## <a href="#_c_specification" class="anchor"></a>C Specification

The list of potential descriptor types a given mutable descriptor
**can** mutate to is passed in a `VkMutableDescriptorTypeListEXT`
structure.

The `VkMutableDescriptorTypeListEXT` structure is defined as:

``` c
// Provided by VK_EXT_mutable_descriptor_type
typedef struct VkMutableDescriptorTypeListEXT {
    uint32_t                   descriptorTypeCount;
    const VkDescriptorType*    pDescriptorTypes;
} VkMutableDescriptorTypeListEXT;
```

or the equivalent

``` c
// Provided by VK_VALVE_mutable_descriptor_type
typedef VkMutableDescriptorTypeListEXT VkMutableDescriptorTypeListVALVE;
```

## <a href="#_members" class="anchor"></a>Members

- `descriptorTypeCount` is the number of elements in `pDescriptorTypes`.

- `pDescriptorTypes` is `NULL` or a pointer to an array of
  `descriptorTypeCount` [VkDescriptorType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorType.html) values
  defining which descriptor types a given binding may mutate to.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkMutableDescriptorTypeListEXT-descriptorTypeCount-04597"
  id="VUID-VkMutableDescriptorTypeListEXT-descriptorTypeCount-04597"></a>
  VUID-VkMutableDescriptorTypeListEXT-descriptorTypeCount-04597  
  `descriptorTypeCount` **must** not be `0` if the corresponding binding
  is of `VK_DESCRIPTOR_TYPE_MUTABLE_EXT`

- <a href="#VUID-VkMutableDescriptorTypeListEXT-pDescriptorTypes-04598"
  id="VUID-VkMutableDescriptorTypeListEXT-pDescriptorTypes-04598"></a>
  VUID-VkMutableDescriptorTypeListEXT-pDescriptorTypes-04598  
  `pDescriptorTypes` **must** be a valid pointer to an array of
  `descriptorTypeCount` valid, unique
  [VkDescriptorType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorType.html) values if the given binding
  is of `VK_DESCRIPTOR_TYPE_MUTABLE_EXT` type

- <a href="#VUID-VkMutableDescriptorTypeListEXT-descriptorTypeCount-04599"
  id="VUID-VkMutableDescriptorTypeListEXT-descriptorTypeCount-04599"></a>
  VUID-VkMutableDescriptorTypeListEXT-descriptorTypeCount-04599  
  `descriptorTypeCount` **must** be `0` if the corresponding binding is
  not of `VK_DESCRIPTOR_TYPE_MUTABLE_EXT`

- <a href="#VUID-VkMutableDescriptorTypeListEXT-pDescriptorTypes-04600"
  id="VUID-VkMutableDescriptorTypeListEXT-pDescriptorTypes-04600"></a>
  VUID-VkMutableDescriptorTypeListEXT-pDescriptorTypes-04600  
  `pDescriptorTypes` **must** not contain
  `VK_DESCRIPTOR_TYPE_MUTABLE_EXT`

- <a href="#VUID-VkMutableDescriptorTypeListEXT-pDescriptorTypes-04601"
  id="VUID-VkMutableDescriptorTypeListEXT-pDescriptorTypes-04601"></a>
  VUID-VkMutableDescriptorTypeListEXT-pDescriptorTypes-04601  
  `pDescriptorTypes` **must** not contain
  `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER_DYNAMIC`

- <a href="#VUID-VkMutableDescriptorTypeListEXT-pDescriptorTypes-04602"
  id="VUID-VkMutableDescriptorTypeListEXT-pDescriptorTypes-04602"></a>
  VUID-VkMutableDescriptorTypeListEXT-pDescriptorTypes-04602  
  `pDescriptorTypes` **must** not contain
  `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER_DYNAMIC`

- <a href="#VUID-VkMutableDescriptorTypeListEXT-pDescriptorTypes-04603"
  id="VUID-VkMutableDescriptorTypeListEXT-pDescriptorTypes-04603"></a>
  VUID-VkMutableDescriptorTypeListEXT-pDescriptorTypes-04603  
  `pDescriptorTypes` **must** not contain
  `VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK`

Valid Usage (Implicit)

- <a
  href="#VUID-VkMutableDescriptorTypeListEXT-pDescriptorTypes-parameter"
  id="VUID-VkMutableDescriptorTypeListEXT-pDescriptorTypes-parameter"></a>
  VUID-VkMutableDescriptorTypeListEXT-pDescriptorTypes-parameter  
  If `descriptorTypeCount` is not `0`, `pDescriptorTypes` **must** be a
  valid pointer to an array of `descriptorTypeCount` valid
  [VkDescriptorType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorType.html) values

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_mutable_descriptor_type](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_mutable_descriptor_type.html),
[VK_VALVE_mutable_descriptor_type](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VALVE_mutable_descriptor_type.html),
[VkDescriptorType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorType.html),
[VkMutableDescriptorTypeCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMutableDescriptorTypeCreateInfoEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkMutableDescriptorTypeListEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
