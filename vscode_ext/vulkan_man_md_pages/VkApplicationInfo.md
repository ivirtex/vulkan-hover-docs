# VkApplicationInfo(3) Manual Page

## Name

VkApplicationInfo - Structure specifying application information



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkApplicationInfo` structure is defined as:

``` c
// Provided by VK_VERSION_1_0
typedef struct VkApplicationInfo {
    VkStructureType    sType;
    const void*        pNext;
    const char*        pApplicationName;
    uint32_t           applicationVersion;
    const char*        pEngineName;
    uint32_t           engineVersion;
    uint32_t           apiVersion;
} VkApplicationInfo;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `pApplicationName` is `NULL` or is a pointer to a null-terminated
  UTF-8 string containing the name of the application.

- `applicationVersion` is an unsigned integer variable containing the
  developer-supplied version number of the application.

- `pEngineName` is `NULL` or is a pointer to a null-terminated UTF-8
  string containing the name of the engine (if any) used to create the
  application.

- `engineVersion` is an unsigned integer variable containing the
  developer-supplied version number of the engine used to create the
  application.

- `apiVersion` **must** be the highest version of Vulkan that the
  application is designed to use, encoded as described in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#extendingvulkan-coreversions-versionnumbers"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#extendingvulkan-coreversions-versionnumbers</a>.
  The patch version number specified in `apiVersion` is ignored when
  creating an instance object. The variant version of the instance
  **must** match that requested in `apiVersion`.

## <a href="#_description" class="anchor"></a>Description

Vulkan 1.0 implementations were required to return
`VK_ERROR_INCOMPATIBLE_DRIVER` if `apiVersion` was larger than 1.0.
Implementations that support Vulkan 1.1 or later **must** not return
`VK_ERROR_INCOMPATIBLE_DRIVER` for any value of `apiVersion` .

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>Because Vulkan 1.0 implementations <strong>may</strong> fail with
<code>VK_ERROR_INCOMPATIBLE_DRIVER</code>, applications
<strong>should</strong> determine the version of Vulkan available before
calling <a href="vkCreateInstance.html">vkCreateInstance</a>. If the <a
href="vkGetInstanceProcAddr.html">vkGetInstanceProcAddr</a> returns
<code>NULL</code> for <a
href="vkEnumerateInstanceVersion.html">vkEnumerateInstanceVersion</a>,
it is a Vulkan 1.0 implementation. Otherwise, the application
<strong>can</strong> call <a
href="vkEnumerateInstanceVersion.html">vkEnumerateInstanceVersion</a> to
determine the version of Vulkan.</p></td>
</tr>
</tbody>
</table>

As long as the instance supports at least Vulkan 1.1, an application
**can** use different versions of Vulkan with an instance than it does
with a device or physical device.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>The Khronos validation layers will treat <code>apiVersion</code> as
the highest API version the application targets, and will validate API
usage against the minimum of that version and the implementation version
(instance or device, depending on context). If an application tries to
use functionality from a greater version than this, a validation error
will be triggered.</p>
<p>For example, if the instance supports Vulkan 1.1 and three physical
devices support Vulkan 1.0, Vulkan 1.1, and Vulkan 1.2, respectively,
and if the application sets <code>apiVersion</code> to 1.2, the
application <strong>can</strong> use the following versions of
Vulkan:</p>
<ul>
<li><p>Vulkan 1.0 <strong>can</strong> be used with the instance and
with all physical devices.</p></li>
<li><p>Vulkan 1.1 <strong>can</strong> be used with the instance and
with the physical devices that support Vulkan 1.1 and Vulkan
1.2.</p></li>
<li><p>Vulkan 1.2 <strong>can</strong> be used with the physical device
that supports Vulkan 1.2.</p></li>
</ul>
<p>If we modify the above example so that the application sets
<code>apiVersion</code> to 1.1, then the application
<strong>must</strong> not use Vulkan 1.2 functionality on the physical
device that supports Vulkan 1.2.</p></td>
</tr>
</tbody>
</table>

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>Providing a <code>NULL</code> <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstanceCreateInfo.html">VkInstanceCreateInfo</a>::<code>pApplicationInfo</code>
or providing an <code>apiVersion</code> of 0 is equivalent to providing
an <code>apiVersion</code> of
<code>VK_MAKE_API_VERSION(0,1,0,0)</code>.</p></td>
</tr>
</tbody>
</table>

Valid Usage

- <a href="#VUID-VkApplicationInfo-apiVersion-04010"
  id="VUID-VkApplicationInfo-apiVersion-04010"></a>
  VUID-VkApplicationInfo-apiVersion-04010  
  If `apiVersion` is not `0`, then it **must** be greater than or equal
  to [VK_API_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_API_VERSION_1_0.html)

Valid Usage (Implicit)

- <a href="#VUID-VkApplicationInfo-sType-sType"
  id="VUID-VkApplicationInfo-sType-sType"></a>
  VUID-VkApplicationInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_APPLICATION_INFO`

- <a href="#VUID-VkApplicationInfo-pNext-pNext"
  id="VUID-VkApplicationInfo-pNext-pNext"></a>
  VUID-VkApplicationInfo-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkApplicationInfo-pApplicationName-parameter"
  id="VUID-VkApplicationInfo-pApplicationName-parameter"></a>
  VUID-VkApplicationInfo-pApplicationName-parameter  
  If `pApplicationName` is not `NULL`, `pApplicationName` **must** be a
  null-terminated UTF-8 string

- <a href="#VUID-VkApplicationInfo-pEngineName-parameter"
  id="VUID-VkApplicationInfo-pEngineName-parameter"></a>
  VUID-VkApplicationInfo-pEngineName-parameter  
  If `pEngineName` is not `NULL`, `pEngineName` **must** be a
  null-terminated UTF-8 string

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkInstanceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstanceCreateInfo.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkApplicationInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
