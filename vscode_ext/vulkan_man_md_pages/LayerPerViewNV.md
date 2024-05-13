# LayerPerViewNV(3) Manual Page

## Name

LayerPerViewNV - Layer index per view for layered rendering



## <a href="#_description" class="anchor"></a>Description

`LayerPerViewNV`  
Decorating a variable with the `LayerPerViewNV` built-in decoration will
make that variable contain the per-view layer information. The per-view
layer has the same semantics as `Layer`, for each view.

Valid Usage

- <a href="#VUID-LayerPerViewNV-LayerPerViewNV-04277"
  id="VUID-LayerPerViewNV-LayerPerViewNV-04277"></a>
  VUID-LayerPerViewNV-LayerPerViewNV-04277  
  The `LayerPerViewNV` decoration **must** be used only within the
  `MeshNV` `Execution` `Model`

- <a href="#VUID-LayerPerViewNV-LayerPerViewNV-04278"
  id="VUID-LayerPerViewNV-LayerPerViewNV-04278"></a>
  VUID-LayerPerViewNV-LayerPerViewNV-04278  
  The variable decorated with `LayerPerViewNV` **must** be declared
  using the `Output` `Storage` `Class`

- <a href="#VUID-LayerPerViewNV-LayerPerViewNV-04279"
  id="VUID-LayerPerViewNV-LayerPerViewNV-04279"></a>
  VUID-LayerPerViewNV-LayerPerViewNV-04279  
  The variable decorated with `LayerPerViewNV` **must** also be
  decorated with the `PerViewNV` decoration

- <a href="#VUID-LayerPerViewNV-LayerPerViewNV-04280"
  id="VUID-LayerPerViewNV-LayerPerViewNV-04280"></a>
  VUID-LayerPerViewNV-LayerPerViewNV-04280  
  The variable decorated with `LayerPerViewNV` **must** be declared as
  an array of scalar 32-bit integer values

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#LayerPerViewNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
