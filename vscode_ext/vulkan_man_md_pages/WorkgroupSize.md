# WorkgroupSize(3) Manual Page

## Name

WorkgroupSize - Size of a workgroup



## <a href="#_description" class="anchor"></a>Description

`WorkgroupSize`

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>SPIR-V 1.6 deprecated <code>WorkgroupSize</code> in favor of using
the <code>LocalSizeId</code> Execution Mode instead. Support for
<code>LocalSizeId</code> was added with <a
href="VK_KHR_maintenance4.html"><code>VK_KHR_maintenance4</code></a> and
promoted to core in Version 1.3.</p></td>
</tr>
</tbody>
</table>

Decorating an object with the `WorkgroupSize` built-in decoration will
make that object contain the dimensions of a local workgroup. If an
object is decorated with the `WorkgroupSize` decoration, this takes
precedence over any `LocalSize` or `LocalSizeId` execution mode.

Valid Usage

- <a href="#VUID-WorkgroupSize-WorkgroupSize-04425"
  id="VUID-WorkgroupSize-WorkgroupSize-04425"></a>
  VUID-WorkgroupSize-WorkgroupSize-04425  
  The `WorkgroupSize` decoration **must** be used only within the
  `GLCompute`, `MeshEXT`, `TaskEXT`, `MeshNV`, or `TaskNV` `Execution`
  `Model`

- <a href="#VUID-WorkgroupSize-WorkgroupSize-04426"
  id="VUID-WorkgroupSize-WorkgroupSize-04426"></a>
  VUID-WorkgroupSize-WorkgroupSize-04426  
  The variable decorated with `WorkgroupSize` **must** be a
  specialization constant or a constant

- <a href="#VUID-WorkgroupSize-WorkgroupSize-04427"
  id="VUID-WorkgroupSize-WorkgroupSize-04427"></a>
  VUID-WorkgroupSize-WorkgroupSize-04427  
  The variable decorated with `WorkgroupSize` **must** be declared as a
  three-component vector of 32-bit integer values

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#WorkgroupSize"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
