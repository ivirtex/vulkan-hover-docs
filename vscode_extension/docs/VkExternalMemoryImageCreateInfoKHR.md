# VkExternalMemoryImageCreateInfo(3) Manual Page

## Name

VkExternalMemoryImageCreateInfo - Specify that an image may be backed by
external memory



## <a href="#_c_specification" class="anchor"></a>C Specification

To define a set of external memory handle types that **may** be used as
backing store for an image, add a
[VkExternalMemoryImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryImageCreateInfo.html)
structure to the `pNext` chain of the
[VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) structure. The
`VkExternalMemoryImageCreateInfo` structure is defined as:

``` c
// Provided by VK_VERSION_1_1
typedef struct VkExternalMemoryImageCreateInfo {
    VkStructureType                    sType;
    const void*                        pNext;
    VkExternalMemoryHandleTypeFlags    handleTypes;
} VkExternalMemoryImageCreateInfo;
```

or the equivalent

``` c
// Provided by VK_KHR_external_memory
typedef VkExternalMemoryImageCreateInfo VkExternalMemoryImageCreateInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>A <code>VkExternalMemoryImageCreateInfo</code> structure with a
non-zero <code>handleTypes</code> field must be included in the creation
parameters for an image that will be bound to memory that is either
exported or imported.</p></td>
</tr>
</tbody>
</table>

## <a href="#_description" class="anchor"></a>Description

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `handleTypes` is zero or a bitmask of
  [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagBits.html)
  specifying one or more external memory handle types.

Valid Usage (Implicit)

- <a href="#VUID-VkExternalMemoryImageCreateInfo-sType-sType"
  id="VUID-VkExternalMemoryImageCreateInfo-sType-sType"></a>
  VUID-VkExternalMemoryImageCreateInfo-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_EXTERNAL_MEMORY_IMAGE_CREATE_INFO`

- <a href="#VUID-VkExternalMemoryImageCreateInfo-handleTypes-parameter"
  id="VUID-VkExternalMemoryImageCreateInfo-handleTypes-parameter"></a>
  VUID-VkExternalMemoryImageCreateInfo-handleTypes-parameter  
  `handleTypes` **must** be a valid combination of
  [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagBits.html)
  values

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html),
[VkExternalMemoryHandleTypeFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlags.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkExternalMemoryImageCreateInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
