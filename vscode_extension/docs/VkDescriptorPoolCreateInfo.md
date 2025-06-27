# VkDescriptorPoolCreateInfo(3) Manual Page

## Name

VkDescriptorPoolCreateInfo - Structure specifying parameters of a newly
created descriptor pool



## <a href="#_c_specification" class="anchor"></a>C Specification

Additional information about the pool is passed in a
`VkDescriptorPoolCreateInfo` structure:

``` c
// Provided by VK_VERSION_1_0
typedef struct VkDescriptorPoolCreateInfo {
    VkStructureType                sType;
    const void*                    pNext;
    VkDescriptorPoolCreateFlags    flags;
    uint32_t                       maxSets;
    uint32_t                       poolSizeCount;
    const VkDescriptorPoolSize*    pPoolSizes;
} VkDescriptorPoolCreateInfo;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is a bitmask of
  [VkDescriptorPoolCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorPoolCreateFlagBits.html)
  specifying certain supported operations on the pool.

- `maxSets` is the maximum number of descriptor sets that **can** be
  allocated from the pool.

- `poolSizeCount` is the number of elements in `pPoolSizes`.

- `pPoolSizes` is a pointer to an array of
  [VkDescriptorPoolSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorPoolSize.html) structures, each
  containing a descriptor type and number of descriptors of that type to
  be allocated in the pool.

## <a href="#_description" class="anchor"></a>Description

If multiple `VkDescriptorPoolSize` structures containing the same
descriptor type appear in the `pPoolSizes` array then the pool will be
created with enough storage for the total number of descriptors of each
type.

Fragmentation of a descriptor pool is possible and **may** lead to
descriptor set allocation failures. A failure due to fragmentation is
defined as failing a descriptor set allocation despite the sum of all
outstanding descriptor set allocations from the pool plus the requested
allocation requiring no more than the total number of descriptors
requested at pool creation. Implementations provide certain guarantees
of when fragmentation **must** not cause allocation failure, as
described below.

If a descriptor pool has not had any descriptor sets freed since it was
created or most recently reset then fragmentation **must** not cause an
allocation failure (note that this is always the case for a pool created
without the `VK_DESCRIPTOR_POOL_CREATE_FREE_DESCRIPTOR_SET_BIT` bit
set). Additionally, if all sets allocated from the pool since it was
created or most recently reset use the same number of descriptors (of
each type) and the requested allocation also uses that same number of
descriptors (of each type), then fragmentation **must** not cause an
allocation failure.

If an allocation failure occurs due to fragmentation, an application
**can** create an additional descriptor pool to perform further
descriptor set allocations.

If `flags` has the `VK_DESCRIPTOR_POOL_CREATE_UPDATE_AFTER_BIND_BIT` bit
set, descriptor pool creation **may** fail with the error
`VK_ERROR_FRAGMENTATION` if the total number of descriptors across all
pools (including this one) created with this bit set exceeds
`maxUpdateAfterBindDescriptorsInAllPools`, or if fragmentation of the
underlying hardware resources occurs.

If a `pPoolSizes`\[i\]::`type` is `VK_DESCRIPTOR_TYPE_MUTABLE_EXT`, a
[VkMutableDescriptorTypeCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMutableDescriptorTypeCreateInfoEXT.html)
struct in the `pNext` chain **can** be used to specify which mutable
descriptor types **can** be allocated from the pool. If included in the
`pNext` chain,
[VkMutableDescriptorTypeCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMutableDescriptorTypeCreateInfoEXT.html)::`pMutableDescriptorTypeLists`\[i\]
specifies which kind of `VK_DESCRIPTOR_TYPE_MUTABLE_EXT` descriptors
**can** be allocated from this pool entry. If
[VkMutableDescriptorTypeCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMutableDescriptorTypeCreateInfoEXT.html)
does not exist in the `pNext` chain, or
[VkMutableDescriptorTypeCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMutableDescriptorTypeCreateInfoEXT.html)::`pMutableDescriptorTypeLists`\[i\]
is out of range, the descriptor pool allocates enough memory to be able
to allocate a `VK_DESCRIPTOR_TYPE_MUTABLE_EXT` descriptor with any
supported [VkDescriptorType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorType.html) as a mutable
descriptor. A mutable descriptor **can** be allocated from a pool entry
if the type list in
[VkDescriptorSetLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayoutCreateInfo.html)
is a subset of the type list declared in the descriptor pool, or if the
pool entry is created without a descriptor type list. Multiple
`pPoolSizes` entries with `VK_DESCRIPTOR_TYPE_MUTABLE_EXT` **can** be
declared. When multiple such pool entries are present in `pPoolSizes`,
they specify sets of supported descriptor types which either fully
overlap, partially overlap, or are disjoint. Two sets fully overlap if
the sets of supported descriptor types are equal. If the sets are not
disjoint they partially overlap. A pool entry without a
`VkMutableDescriptorTypeListEXT` assigned to it is considered to
partially overlap any other pool entry which has a
`VkMutableDescriptorTypeListEXT` assigned to it. The application
**must** ensure that partial overlap does not exist in `pPoolSizes`.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>The requirement of no partial overlap is intended to resolve
ambiguity for validation as there is no confusion which
<code>pPoolSizes</code> entries will be allocated from. An
implementation is not expected to depend on this requirement.</p></td>
</tr>
</tbody>
</table>

Valid Usage

- <a
  href="#VUID-VkDescriptorPoolCreateInfo-descriptorPoolOverallocation-09227"
  id="VUID-VkDescriptorPoolCreateInfo-descriptorPoolOverallocation-09227"></a>
  VUID-VkDescriptorPoolCreateInfo-descriptorPoolOverallocation-09227  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-descriptorPoolOverallocation"
  target="_blank"
  rel="noopener"><code>descriptorPoolOverallocation</code></a> feature
  is not enabled, or `flags` does not have
  `VK_DESCRIPTOR_POOL_CREATE_ALLOW_OVERALLOCATION_SETS_BIT_NV` set,
  `maxSets` **must** be greater than `0`

- <a href="#VUID-VkDescriptorPoolCreateInfo-flags-09228"
  id="VUID-VkDescriptorPoolCreateInfo-flags-09228"></a>
  VUID-VkDescriptorPoolCreateInfo-flags-09228  
  If `flags` has the
  `VK_DESCRIPTOR_POOL_CREATE_ALLOW_OVERALLOCATION_SETS_BIT_NV` or
  `VK_DESCRIPTOR_POOL_CREATE_ALLOW_OVERALLOCATION_POOLS_BIT_NV` bits
  set, then <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-descriptorPoolOverallocation"
  target="_blank"
  rel="noopener"><code>descriptorPoolOverallocation</code></a> **must**
  be enabled

- <a href="#VUID-VkDescriptorPoolCreateInfo-flags-04607"
  id="VUID-VkDescriptorPoolCreateInfo-flags-04607"></a>
  VUID-VkDescriptorPoolCreateInfo-flags-04607  
  If `flags` has the `VK_DESCRIPTOR_POOL_CREATE_HOST_ONLY_BIT_EXT` bit
  set, then the `VK_DESCRIPTOR_POOL_CREATE_UPDATE_AFTER_BIND_BIT` bit
  **must** not be set

- <a href="#VUID-VkDescriptorPoolCreateInfo-mutableDescriptorType-04608"
  id="VUID-VkDescriptorPoolCreateInfo-mutableDescriptorType-04608"></a>
  VUID-VkDescriptorPoolCreateInfo-mutableDescriptorType-04608  
  If
  [VkPhysicalDeviceMutableDescriptorTypeFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMutableDescriptorTypeFeaturesEXT.html)::`mutableDescriptorType`
  is not enabled, `pPoolSizes` **must** not contain a `descriptorType`
  of `VK_DESCRIPTOR_TYPE_MUTABLE_EXT`

- <a href="#VUID-VkDescriptorPoolCreateInfo-flags-04609"
  id="VUID-VkDescriptorPoolCreateInfo-flags-04609"></a>
  VUID-VkDescriptorPoolCreateInfo-flags-04609  
  If `flags` has the `VK_DESCRIPTOR_POOL_CREATE_HOST_ONLY_BIT_EXT` bit
  set,
  [VkPhysicalDeviceMutableDescriptorTypeFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMutableDescriptorTypeFeaturesEXT.html)::`mutableDescriptorType`
  **must** be enabled

- <a href="#VUID-VkDescriptorPoolCreateInfo-pPoolSizes-04787"
  id="VUID-VkDescriptorPoolCreateInfo-pPoolSizes-04787"></a>
  VUID-VkDescriptorPoolCreateInfo-pPoolSizes-04787  
  If `pPoolSizes` contains a `descriptorType` of
  `VK_DESCRIPTOR_TYPE_MUTABLE_EXT`, any other
  `VK_DESCRIPTOR_TYPE_MUTABLE_EXT` element in `pPoolSizes` **must** not
  have sets of supported descriptor types which partially overlap

- <a href="#VUID-VkDescriptorPoolCreateInfo-pPoolSizes-09424"
  id="VUID-VkDescriptorPoolCreateInfo-pPoolSizes-09424"></a>
  VUID-VkDescriptorPoolCreateInfo-pPoolSizes-09424  
  If `pPoolSizes` contains a `descriptorType` of
  `VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK`, the `pNext` chain **must**
  include a
  [VkDescriptorPoolInlineUniformBlockCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorPoolInlineUniformBlockCreateInfo.html)
  structure whose `maxInlineUniformBlockBindings` member is not zero

Valid Usage (Implicit)

- <a href="#VUID-VkDescriptorPoolCreateInfo-sType-sType"
  id="VUID-VkDescriptorPoolCreateInfo-sType-sType"></a>
  VUID-VkDescriptorPoolCreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DESCRIPTOR_POOL_CREATE_INFO`

- <a href="#VUID-VkDescriptorPoolCreateInfo-pNext-pNext"
  id="VUID-VkDescriptorPoolCreateInfo-pNext-pNext"></a>
  VUID-VkDescriptorPoolCreateInfo-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the
  `pNext` chain **must** be either `NULL` or a pointer to a valid
  instance of
  [VkDescriptorPoolInlineUniformBlockCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorPoolInlineUniformBlockCreateInfo.html)
  or
  [VkMutableDescriptorTypeCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMutableDescriptorTypeCreateInfoEXT.html)

- <a href="#VUID-VkDescriptorPoolCreateInfo-sType-unique"
  id="VUID-VkDescriptorPoolCreateInfo-sType-unique"></a>
  VUID-VkDescriptorPoolCreateInfo-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

- <a href="#VUID-VkDescriptorPoolCreateInfo-flags-parameter"
  id="VUID-VkDescriptorPoolCreateInfo-flags-parameter"></a>
  VUID-VkDescriptorPoolCreateInfo-flags-parameter  
  `flags` **must** be a valid combination of
  [VkDescriptorPoolCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorPoolCreateFlagBits.html)
  values

- <a href="#VUID-VkDescriptorPoolCreateInfo-pPoolSizes-parameter"
  id="VUID-VkDescriptorPoolCreateInfo-pPoolSizes-parameter"></a>
  VUID-VkDescriptorPoolCreateInfo-pPoolSizes-parameter  
  If `poolSizeCount` is not `0`, `pPoolSizes` **must** be a valid
  pointer to an array of `poolSizeCount` valid
  [VkDescriptorPoolSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorPoolSize.html) structures

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkDescriptorPoolCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorPoolCreateFlags.html),
[VkDescriptorPoolSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorPoolSize.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCreateDescriptorPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateDescriptorPool.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDescriptorPoolCreateInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
