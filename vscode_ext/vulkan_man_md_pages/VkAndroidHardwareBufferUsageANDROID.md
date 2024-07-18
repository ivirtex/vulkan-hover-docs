# VkAndroidHardwareBufferUsageANDROID(3) Manual Page

## Name

VkAndroidHardwareBufferUsageANDROID - Struct containing Android hardware
buffer usage flags



## <a href="#_c_specification" class="anchor"></a>C Specification

To obtain optimal Android hardware buffer usage flags for specific image
creation parameters, add a `VkAndroidHardwareBufferUsageANDROID`
structure to the `pNext` chain of a
[VkImageFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageFormatProperties2.html) structure
passed to
[vkGetPhysicalDeviceImageFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceImageFormatProperties2.html).
This structure is defined as:

``` c
// Provided by VK_ANDROID_external_memory_android_hardware_buffer
typedef struct VkAndroidHardwareBufferUsageANDROID {
    VkStructureType    sType;
    void*              pNext;
    uint64_t           androidHardwareBufferUsage;
} VkAndroidHardwareBufferUsageANDROID;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `androidHardwareBufferUsage` returns the Android hardware buffer usage
  flags.

## <a href="#_description" class="anchor"></a>Description

The `androidHardwareBufferUsage` field **must** include Android hardware
buffer usage flags listed in the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-external-android-hardware-buffer-usage"
target="_blank" rel="noopener">AHardwareBuffer Usage Equivalence</a>
table when the corresponding Vulkan image usage or image creation flags
are included in the `usage` or `flags` fields of
[VkPhysicalDeviceImageFormatInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageFormatInfo2.html).
It **must** include at least one GPU usage flag
(`AHARDWAREBUFFER_USAGE_GPU_*`), even if none of the corresponding
Vulkan usages or flags are requested.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>Requiring at least one GPU usage flag ensures that Android hardware
buffer memory will be allocated in a memory pool accessible to the
Vulkan implementation, and that specializing the memory layout based on
usage flags does not prevent it from being compatible with Vulkan.
Implementations <strong>may</strong> avoid unnecessary restrictions
caused by this requirement by using vendor usage flags to indicate that
only the Vulkan uses indicated in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageFormatProperties2.html">VkImageFormatProperties2</a> are
required.</p></td>
</tr>
</tbody>
</table>

Valid Usage (Implicit)

- <a href="#VUID-VkAndroidHardwareBufferUsageANDROID-sType-sType"
  id="VUID-VkAndroidHardwareBufferUsageANDROID-sType-sType"></a>
  VUID-VkAndroidHardwareBufferUsageANDROID-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_ANDROID_HARDWARE_BUFFER_USAGE_ANDROID`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_ANDROID_external_memory_android_hardware_buffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_ANDROID_external_memory_android_hardware_buffer.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkAndroidHardwareBufferUsageANDROID"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
