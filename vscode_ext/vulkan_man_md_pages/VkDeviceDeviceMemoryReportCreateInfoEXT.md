# VkDeviceDeviceMemoryReportCreateInfoEXT(3) Manual Page

## Name

VkDeviceDeviceMemoryReportCreateInfoEXT - Register device memory report
callbacks for a Vulkan device



## <a href="#_c_specification" class="anchor"></a>C Specification

To register callbacks for underlying device memory events of type
[VkDeviceMemoryReportEventTypeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemoryReportEventTypeEXT.html),
add one or multiple
[VkDeviceDeviceMemoryReportCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceDeviceMemoryReportCreateInfoEXT.html)
structures to the `pNext` chain of the
[VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) structure.

``` c
// Provided by VK_EXT_device_memory_report
typedef struct VkDeviceDeviceMemoryReportCreateInfoEXT {
    VkStructureType                        sType;
    const void*                            pNext;
    VkDeviceMemoryReportFlagsEXT           flags;
    PFN_vkDeviceMemoryReportCallbackEXT    pfnUserCallback;
    void*                                  pUserData;
} VkDeviceDeviceMemoryReportCreateInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is 0 and reserved for future use.

- `pfnUserCallback` is the application callback function to call.

- `pUserData` is user data to be passed to the callback.

## <a href="#_description" class="anchor"></a>Description

The callback **may** be called from multiple threads simultaneously.

The callback **must** be called only once by the implementation when a
[VkDeviceMemoryReportEventTypeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemoryReportEventTypeEXT.html)
event occurs.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>The callback could be called from a background thread other than the
thread calling the Vulkan commands.</p></td>
</tr>
</tbody>
</table>

Valid Usage (Implicit)

- <a href="#VUID-VkDeviceDeviceMemoryReportCreateInfoEXT-sType-sType"
  id="VUID-VkDeviceDeviceMemoryReportCreateInfoEXT-sType-sType"></a>
  VUID-VkDeviceDeviceMemoryReportCreateInfoEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_DEVICE_DEVICE_MEMORY_REPORT_CREATE_INFO_EXT`

- <a
  href="#VUID-VkDeviceDeviceMemoryReportCreateInfoEXT-flags-zerobitmask"
  id="VUID-VkDeviceDeviceMemoryReportCreateInfoEXT-flags-zerobitmask"></a>
  VUID-VkDeviceDeviceMemoryReportCreateInfoEXT-flags-zerobitmask  
  `flags` **must** be `0`

- <a
  href="#VUID-VkDeviceDeviceMemoryReportCreateInfoEXT-pfnUserCallback-parameter"
  id="VUID-VkDeviceDeviceMemoryReportCreateInfoEXT-pfnUserCallback-parameter"></a>
  VUID-VkDeviceDeviceMemoryReportCreateInfoEXT-pfnUserCallback-parameter  
  `pfnUserCallback` **must** be a valid
  [PFN_vkDeviceMemoryReportCallbackEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/PFN_vkDeviceMemoryReportCallbackEXT.html)
  value

- <a
  href="#VUID-VkDeviceDeviceMemoryReportCreateInfoEXT-pUserData-parameter"
  id="VUID-VkDeviceDeviceMemoryReportCreateInfoEXT-pUserData-parameter"></a>
  VUID-VkDeviceDeviceMemoryReportCreateInfoEXT-pUserData-parameter  
  `pUserData` **must** be a pointer value

## <a href="#_see_also" class="anchor"></a>See Also

[PFN_vkDeviceMemoryReportCallbackEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/PFN_vkDeviceMemoryReportCallbackEXT.html),
[VK_EXT_device_memory_report](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_device_memory_report.html),
[VkDeviceMemoryReportFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemoryReportFlagsEXT.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDeviceDeviceMemoryReportCreateInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
