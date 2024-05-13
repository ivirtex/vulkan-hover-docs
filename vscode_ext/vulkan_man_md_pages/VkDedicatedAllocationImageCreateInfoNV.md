# VkDedicatedAllocationImageCreateInfoNV(3) Manual Page

## Name

VkDedicatedAllocationImageCreateInfoNV - Specify that an image is bound
to a dedicated memory resource



## <a href="#_c_specification" class="anchor"></a>C Specification

If the `pNext` chain includes a `VkDedicatedAllocationImageCreateInfoNV`
structure, then that structure includes an enable controlling whether
the image will have a dedicated memory allocation bound to it.

The `VkDedicatedAllocationImageCreateInfoNV` structure is defined as:

``` c
// Provided by VK_NV_dedicated_allocation
typedef struct VkDedicatedAllocationImageCreateInfoNV {
    VkStructureType    sType;
    const void*        pNext;
    VkBool32           dedicatedAllocation;
} VkDedicatedAllocationImageCreateInfoNV;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `dedicatedAllocation` specifies whether the image will have a
  dedicated allocation bound to it.

## <a href="#_description" class="anchor"></a>Description

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>Using a dedicated allocation for color and depth/stencil attachments
or other large images <strong>may</strong> improve performance on some
devices.</p></td>
</tr>
</tbody>
</table>

Valid Usage

- <a
  href="#VUID-VkDedicatedAllocationImageCreateInfoNV-dedicatedAllocation-00994"
  id="VUID-VkDedicatedAllocationImageCreateInfoNV-dedicatedAllocation-00994"></a>
  VUID-VkDedicatedAllocationImageCreateInfoNV-dedicatedAllocation-00994  
  If `dedicatedAllocation` is `VK_TRUE`,
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`flags` **must** not
  include `VK_IMAGE_CREATE_SPARSE_BINDING_BIT`,
  `VK_IMAGE_CREATE_SPARSE_RESIDENCY_BIT`, or
  `VK_IMAGE_CREATE_SPARSE_ALIASED_BIT`

Valid Usage (Implicit)

- <a href="#VUID-VkDedicatedAllocationImageCreateInfoNV-sType-sType"
  id="VUID-VkDedicatedAllocationImageCreateInfoNV-sType-sType"></a>
  VUID-VkDedicatedAllocationImageCreateInfoNV-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_DEDICATED_ALLOCATION_IMAGE_CREATE_INFO_NV`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_dedicated_allocation](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_dedicated_allocation.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDedicatedAllocationImageCreateInfoNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
