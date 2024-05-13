# VkDeviceQueueInfo2(3) Manual Page

## Name

VkDeviceQueueInfo2 - Structure specifying the parameters used for device
queue creation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkDeviceQueueInfo2` structure is defined as:

``` c
// Provided by VK_VERSION_1_1
typedef struct VkDeviceQueueInfo2 {
    VkStructureType             sType;
    const void*                 pNext;
    VkDeviceQueueCreateFlags    flags;
    uint32_t                    queueFamilyIndex;
    uint32_t                    queueIndex;
} VkDeviceQueueInfo2;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure. The `pNext` chain of `VkDeviceQueueInfo2` **can** be used
  to provide additional device queue parameters to `vkGetDeviceQueue2`.

- `flags` is a [VkDeviceQueueCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceQueueCreateFlags.html)
  value indicating the flags used to create the device queue.

- `queueFamilyIndex` is the index of the queue family to which the queue
  belongs.

- `queueIndex` is the index of the queue to retrieve from within the set
  of queues that share both the queue family and flags specified.

## <a href="#_description" class="anchor"></a>Description

The queue returned by `vkGetDeviceQueue2` **must** have the same `flags`
value from this structure as that used at device creation time in a
[VkDeviceQueueCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceQueueCreateInfo.html) structure.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>Normally, if you create both protected-capable and
non-protected-capable queues with the same family, they are treated as
separate lists of queues and <code>queueIndex</code> is relative to the
start of the list of queues specified by both
<code>queueFamilyIndex</code> and <code>flags</code>. However, for
historical reasons, some implementations may exhibit different behavior.
These divergent implementations instead concatenate the lists of queues
and treat <code>queueIndex</code> as relative to the start of the first
list of queues with the given <code>queueFamilyIndex</code>. This only
matters in cases where an application has created both protected-capable
and non-protected-capable queues from the same queue family.</p>
<p>For such divergent implementations, the maximum value of
<code>queueIndex</code> is equal to the sum of <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceQueueCreateInfo.html">VkDeviceQueueCreateInfo</a>::<code>queueCount</code>
minus one, for all <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceQueueCreateInfo.html">VkDeviceQueueCreateInfo</a>
structures that share a common <code>queueFamilyIndex</code>.</p>
<p>Such implementations will return <code>NULL</code> for either the
protected or unprotected queues when calling
<code>vkGetDeviceQueue2</code> with <code>queueIndex</code> in the range
zero to <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceQueueCreateInfo.html">VkDeviceQueueCreateInfo</a>::<code>queueCount</code>
minus one. In cases where these implementations returned
<code>NULL</code>, the corresponding queues are instead located in the
extended range described in the preceding two paragraphs.</p>
<p>This behavior will not be observed on any driver that has passed
Vulkan conformance test suite version 1.3.3.0, or any subsequent
version. This information can be found by querying
<code>VkPhysicalDeviceDriverProperties</code>::<code>conformanceVersion</code>.</p></td>
</tr>
</tbody>
</table>

Valid Usage

- <a href="#VUID-VkDeviceQueueInfo2-queueFamilyIndex-01842"
  id="VUID-VkDeviceQueueInfo2-queueFamilyIndex-01842"></a>
  VUID-VkDeviceQueueInfo2-queueFamilyIndex-01842  
  `queueFamilyIndex` **must** be one of the queue family indices
  specified when `device` was created, via the
  [VkDeviceQueueCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceQueueCreateInfo.html) structure

- <a href="#VUID-VkDeviceQueueInfo2-flags-06225"
  id="VUID-VkDeviceQueueInfo2-flags-06225"></a>
  VUID-VkDeviceQueueInfo2-flags-06225  
  `flags` **must** be equal to
  [VkDeviceQueueCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceQueueCreateInfo.html)::`flags` for a
  [VkDeviceQueueCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceQueueCreateInfo.html) structure for
  the queue family indicated by `queueFamilyIndex` when `device` was
  created

- <a href="#VUID-VkDeviceQueueInfo2-queueIndex-01843"
  id="VUID-VkDeviceQueueInfo2-queueIndex-01843"></a>
  VUID-VkDeviceQueueInfo2-queueIndex-01843  
  `queueIndex` **must** be less than
  [VkDeviceQueueCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceQueueCreateInfo.html)::`queueCount`
  for the corresponding queue family and flags indicated by
  `queueFamilyIndex` and `flags` when `device` was created

Valid Usage (Implicit)

- <a href="#VUID-VkDeviceQueueInfo2-sType-sType"
  id="VUID-VkDeviceQueueInfo2-sType-sType"></a>
  VUID-VkDeviceQueueInfo2-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DEVICE_QUEUE_INFO_2`

- <a href="#VUID-VkDeviceQueueInfo2-pNext-pNext"
  id="VUID-VkDeviceQueueInfo2-pNext-pNext"></a>
  VUID-VkDeviceQueueInfo2-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkDeviceQueueInfo2-flags-parameter"
  id="VUID-VkDeviceQueueInfo2-flags-parameter"></a>
  VUID-VkDeviceQueueInfo2-flags-parameter  
  `flags` **must** be a valid combination of
  [VkDeviceQueueCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceQueueCreateFlagBits.html) values

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html),
[VkDeviceQueueCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceQueueCreateFlags.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetDeviceQueue2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeviceQueue2.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDeviceQueueInfo2"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
