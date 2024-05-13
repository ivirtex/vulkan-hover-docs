# vkEnumerateInstanceVersion(3) Manual Page

## Name

vkEnumerateInstanceVersion - Query instance-level version before
instance creation



## <a href="#_c_specification" class="anchor"></a>C Specification

To query the version of instance-level functionality supported by the
implementation, call:

``` c
// Provided by VK_VERSION_1_1
VkResult vkEnumerateInstanceVersion(
    uint32_t*                                   pApiVersion);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `pApiVersion` is a pointer to a `uint32_t`, which is the version of
  Vulkan supported by instance-level functionality, encoded as described
  in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#extendingvulkan-coreversions-versionnumbers"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#extendingvulkan-coreversions-versionnumbers</a>.

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
<p>The intended behavior of <a
href="vkEnumerateInstanceVersion.html">vkEnumerateInstanceVersion</a> is
that an implementation <strong>should</strong> not need to perform
memory allocations and <strong>should</strong> unconditionally return
<code>VK_SUCCESS</code>. The loader, and any enabled layers,
<strong>may</strong> return <code>VK_ERROR_OUT_OF_HOST_MEMORY</code> in
the case of a failed memory allocation.</p></td>
</tr>
</tbody>
</table>

Valid Usage (Implicit)

- <a href="#VUID-vkEnumerateInstanceVersion-pApiVersion-parameter"
  id="VUID-vkEnumerateInstanceVersion-pApiVersion-parameter"></a>
  VUID-vkEnumerateInstanceVersion-pApiVersion-parameter  
  `pApiVersion` **must** be a valid pointer to a `uint32_t` value

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkEnumerateInstanceVersion"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
