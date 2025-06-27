# PFN_vkGetInstanceProcAddrLUNARG(3) Manual Page

## Name

PFN_vkGetInstanceProcAddrLUNARG - Type definition for
vkGetInstanceProcAddr



## <a href="#_c_specification" class="anchor"></a>C Specification

The type of
[PFN_vkGetInstanceProcAddrLUNARG](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/PFN_vkGetInstanceProcAddrLUNARG.html)
is:

``` c
// Provided by VK_LUNARG_direct_driver_loading
typedef PFN_vkVoidFunction (VKAPI_PTR *PFN_vkGetInstanceProcAddrLUNARG)(
    VkInstance instance, const char* pName);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `instance` is a [VkInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstance.html) handle.

- `pName` is the name of a Vulkan command.

## <a href="#_description" class="anchor"></a>Description

This type is compatible with the type of a pointer to the
[vkGetInstanceProcAddr](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetInstanceProcAddr.html) command, but is used
only to specify device driver addresses in
[VkDirectDriverLoadingInfoLUNARG](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDirectDriverLoadingInfoLUNARG.html)::`pfnGetInstanceProcAddr`.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>This type exists only because of limitations in the XML schema and
processing scripts, and its name may change in the future. Ideally we
would use the <code>PFN_vkGetInstanceProcAddr</code> type generated in
the <code>vulkan_core.h</code> header.</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_LUNARG_direct_driver_loading](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_LUNARG_direct_driver_loading.html),
[VkDirectDriverLoadingInfoLUNARG](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDirectDriverLoadingInfoLUNARG.html),
[VkInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstance.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#PFN_vkGetInstanceProcAddrLUNARG"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
